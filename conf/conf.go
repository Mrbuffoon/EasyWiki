package conf

import "github.com/widuu/goini"

func GetValue(section, key string) string {
	conf := goini.SetConfig("easywiki.ini")
	return conf.GetValue(section, key)
}