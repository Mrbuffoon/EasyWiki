package mdtohtml

import (
	"EasyWiki/log"
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"gopkg.in/russross/blackfriday.v2"
)

type MK struct {
	Content template.HTML
}

/* Markdown文件转换成同名HTML文件 */
func MarkdownToHtml(filepath, templateHtml string) error {
	destFile := strings.ReplaceAll(filepath, ".md", ".html")

	mdStr, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Error.Println("Open Markdown file failed")
		return err
	}

	htmlStr := blackfriday.Run(mdStr)
	htmlStr = bytes.ReplaceAll(htmlStr, []byte(".md"), []byte(".html"))
	if strings.HasSuffix(filepath, "RESUME.md") {
		regex := regexp.MustCompile(`<a (href=\".+.html\")>`)
		params := regex.FindAllSubmatch(htmlStr, -1)
		for _, param := range params {
			htmlStr = bytes.ReplaceAll(htmlStr, param[1], []byte(string(param[1]) + " target=\"myiFrame\""))
		}
	}
	content := template.HTML(htmlStr)
	mk := MK{Content: content}

	temp, _ := template.ParseFiles(templateHtml)
	writer, err := os.Create(destFile)
	if err != nil {
		log.Error.Println("Open template file error")
		return err
	}

	return temp.Execute(writer, mk)
}
