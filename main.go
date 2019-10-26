package main

import (
	"EasyWiki/mdtohtml"
	"EasyWiki/conf"
	"EasyWiki/fileops"
	"EasyWiki/log"
)

func main() {
	mdFile := conf.GetValue("SYSTEM", "MdPath")
	htmlPath := conf.GetValue("SYSTEM", "HtmlPath")
	fileops.CopyDir(mdFile)
	mdtohtml.MarkdownToHtml("blogs/blog/代码管理中常用git操作.md")
}
