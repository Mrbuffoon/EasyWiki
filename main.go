package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/Mrbuffoon/EasyWiki/conf"
	"github.com/Mrbuffoon/EasyWiki/fileops"
	"github.com/Mrbuffoon/EasyWiki/log"
	"github.com/Mrbuffoon/EasyWiki/mdtohtml"
)

func PublishWiki() error {
	mdFile := conf.GetValue("SYSTEM", "MdPath")
	if !fileops.FileIsExisted(mdFile) {
		err := fileops.MakeDir(mdFile)
		if err != nil {
			log.Error.Println(err)
			return err
		}
	}

	targetMdFiles, err := fileops.WalkDir(mdFile, ".md")
	if err != nil {
		log.Error.Println(err)
		return err
	}
	for _, targetMdFile := range targetMdFiles {
		if strings.HasSuffix(targetMdFile, "RESUME.md") {
			err = mdtohtml.MarkdownToHtml(targetMdFile, "/var/easywiki/template/easywiki_resume.html")
		} else {
			err = mdtohtml.MarkdownToHtml(targetMdFile, "/var/easywiki/template/easywiki_article.html")
		}
		if err != nil {
			log.Error.Println(err)
			return err
		}
	}

	desDir := conf.GetValue("WEB", "WebRoot")
	if !fileops.FileIsExisted(desDir) {
		err := fileops.MakeDir(desDir)
		if err != nil {
			log.Error.Println(err)
			return err
		}
	}
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

	if !fileops.FileIsExisted(blogPath) {
		err := fileops.MakeDir(blogPath)
		if err != nil {
			log.Error.Println(err)
			return err
		}
	} else {
		err := fileops.RemoveContents(blogPath)
		if err != nil {
			log.Error.Println(err)
			return err
		}
	}

	cmdStr := "git clone " + repoAddr + " " + blogPath
	cmd := exec.Command("bash", "-c", cmdStr)
	err := cmd.Run()
	if err != nil {
		log.Error.Println(err)
		return err
	}

	return nil
}

func runHandler(w http.ResponseWriter, r *http.Request) {
	result, _ := json.Marshal("successful")
	fmt.Fprint(w, string(result))

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
	port := conf.GetValue("PORT", "Port")
	if port == "" {
		port = "9090"
	}
	log.Error.Fatal(http.ListenAndServe(":"+port, nil))
}
