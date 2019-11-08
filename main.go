package main

import (
	"EasyWiki/conf"
	"EasyWiki/fileops"
	"EasyWiki/log"
	"EasyWiki/mdtohtml"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

/*
TODO:
写一个Makefile
配置webhooks
写一个Readme
*/

func PublishWiki() error {
	mdFile := conf.GetValue("SYSTEM", "MdPath")

	targetMdFiles, err := fileops.WalkDir(mdFile, ".md")
	if err != nil {
		log.Error.Println(err)
		return err
	}
	for _, targetMdFile := range targetMdFiles {
		if strings.HasSuffix(targetMdFile, "RESUME.md") {
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
	err = fileops.CopyDir(mdFile, desDir+"/easywikis")
	if err != nil {
		log.Error.Println(err)
		return err
	}

	targetMdFiles, err = fileops.WalkDir(desDir+"/easywikis", ".md")
	if err != nil {
		log.Error.Println(err)
		return err
	}
	for _, targetMdFile := range targetMdFiles {
		if srcInfo, err := os.Stat(targetMdFile); err == nil {
			if !srcInfo.IsDir() {
				os.Remove(targetMdFile)
			}
		}
	}

	_, err = fileops.CopyFile(desDir+"/easywikis/RESUME.html", desDir+"/easywikis/index.html")
	if err != nil {
		log.Error.Println(err)
		return err
	}

	return nil
}

func PullCode() error {
	blogPath := conf.GetValue("SYSTEM", "MdPath")
	repoAddr := conf.GetValue("GIT", "RepoAddr")
	err := fileops.RemoveContents(blogPath)
	if err != nil {
		log.Error.Println(err)
		return err
	}
	cmdStr := "git clone " + repoAddr + " " + blogPath
	cmd := exec.Command("bash", "-c", cmdStr)
	err = cmd.Run()
	if err != nil {
		log.Error.Println(err)
		return err
	}

	return nil
}

func runHandler(w http.ResponseWriter, r *http.Request) {
	err := PullCode()
	if err != nil {
		log.Error.Println("Clone code fail")
	}
	err = PublishWiki()
	if err != nil {
		log.Error.Println("Wiki发布失败，请查看log获取更多信息")
	} else {
		log.Info.Println("Wiki发布成功")
	}
}

func main() {
	http.HandleFunc("/webhooks", runHandler)
	log.Error.Fatal(http.ListenAndServe(":9090", nil))
}
