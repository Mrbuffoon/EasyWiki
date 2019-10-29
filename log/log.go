package log

import (
	"EasyWiki/conf"
	"io"
	"log"
	"os"
)

var (
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	filePath := conf.GetValue("LOG", "LogPath")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err := os.MkdirAll(filePath, 0777); err != nil { //os.ModePerm
			log.Println("创建日志路径失败:", err)
		}
	}
	logFile, err := os.OpenFile(filePath+"/easywiki.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Panic("创建日志文件出错，请检查")
	}
	defer logFile.Close()

	Debug = log.New(os.Stdout, "Debug:", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(io.MultiWriter(os.Stderr, logFile), "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(io.MultiWriter(os.Stderr, logFile), "Warning:", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, logFile), "Error:", log.Ldate|log.Ltime|log.Lshortfile)
}
