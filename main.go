package main

import (
	"EasyWiki/conf"
	"EasyWiki/fileops"
	"EasyWiki/log"
	"EasyWiki/mdtohtml"
)

func main() {
	log.Info.Println("Test Start")
	mdFile := conf.GetValue("SYSTEM", "MdPath")
	htmlPath := conf.GetValue("SYSTEM", "HtmlPath")
	err := fileops.CopyDir(mdFile, htmlPath)
	if err != nil {
		log.Error.Println(err)
		return
	}
	targetMdFiles, err := fileops.WalkDir(htmlPath, ".md")
	if err != nil {
		log.Error.Println(err)
		return
	}
	for _, targetMdFile := range targetMdFiles {
		mdtohtml.MarkdownToHtml(targetMdFile)
	}
}
