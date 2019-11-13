package conf

import (
	"log"
	"os"

	"github.com/widuu/goini"
)

func GetValue(section, key string) string {
	if _, err := os.Stat("/etc/easywiki/easywiki.ini"); os.IsNotExist(err) {
		log.Panic(err)
	}
	conf := goini.SetConfig("/etc/easywiki/easywiki.ini")
	return conf.GetValue(section, key)
}
