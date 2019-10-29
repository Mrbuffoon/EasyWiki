package main

import (
	"EasyWiki/conf"
	"EasyWiki/fileops"
	"EasyWiki/log"
	"EasyWiki/mdtohtml"
	"strings"
)

func PublishWiki() error {
	mdFile := conf.GetValue("SYSTEM", "MdPath")
	htmlPath := conf.GetValue("SYSTEM", "HtmlPath")
	err := fileops.CopyDir(mdFile, htmlPath)
	if err != nil {
		log.Error.Println(err)
		return err
	}

	targetMdFiles, err := fileops.WalkDir(htmlPath, ".md")
	if err != nil {
		log.Error.Println(err)
		return err
	}
	for _, targetMdFile := range targetMdFiles {
		if strings.HasSuffix(targetMdFile, "RESUME.md")  {
			err = mdtohtml.MarkdownToHtml(targetMdFile, "./template/resume.html")
		} else {
			err = mdtohtml.MarkdownToHtml(targetMdFile, "./template/article.html")
		}
		if err != nil {
			log.Error.Println(err)
			return err
		}
	}

	desDir := conf.GetValue("WEB", "WebRoot")
	err = fileops.CopyDir(htmlPath, desDir+"/easywikis")
	if err != nil {
		log.Error.Println(err)
		return err
	}

	return nil
}

func main() {
	err := PublishWiki()
	if err != nil {
		log.Error.Println("Wiki发布失败，请查看log获取更多信息")
	} else {
		log.Info.Println("Wiki发布成功")
	}
}
