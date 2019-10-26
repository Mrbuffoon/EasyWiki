package mdtohtml

import (
	"EasyWiki/log"
	"html/template"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/russross/blackfriday.v2"
)

type MK struct {
	Content template.HTML
}

/* Markdown文件转换成同名HTML文件 */
func MarkdownToHtml(filepath string) error {
	destFile := "./views/" + strings.ReplaceAll(filepath, ".md", ".html")

	mdStr, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Error.Println("Open Markdown file failed")
		return err
	}

	content := template.HTML(blackfriday.Run(mdStr))
	mk := MK{Content: content}

	temp, _ := template.ParseFiles("./template/html/blog.html")
	writer, err := os.Create(destFile)
	if err != nil {
		log.Error.Println("Open template file error")
		return err
	}

	return temp.Execute(writer, mk)
}
