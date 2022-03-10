package tool

import (
	"log"
	"os"
	"strings"
	"unicode"
)

// NameStyle 检查命名规范
func NameStyle(name string, dirPth string) {
	if name == "" {
		log.Fatalln("🐻 大熊弟，项目名称不能为空哟～")
	}

	// 检查开头字母
	for _, n := range name[:1] {
		if unicode.IsUpper(n) {
			log.Fatalln("🐻 大熊弟，项目名称首字母不能为大写～")
		}
		if unicode.IsNumber(n) {
			log.Fatalln("🐻 大熊弟，项目名称不能以数字开头～")
		}
		if unicode.IsSpace(n) {
			log.Fatalln("🐻 大熊弟，项目名称不能以空白开头～")
		}
		if unicode.IsPunct(n) {
			log.Fatalln("🐻 大熊弟，项目名称不能以标点符号开头～")
		}
	}

	// 检查名称字符是否含有特殊字符
	speSth := []string{
		"_",
		"/",
		"-",
		" ",
	}
	for _, s := range speSth {
		if strings.Contains(name, s) {
			log.Fatalln("🐻 大熊弟，项目名称不能有特殊字符～")
		}
	}
	for _, v := range name {
		if unicode.IsPunct(v) {
			log.Fatalln("🐻 大熊弟，项目名称不能含有标点符号～")
		}
	}

	// 判断文件或文件夹是否存在
	_, err := os.Stat(dirPth)
	if err != nil {
		log.Fatalln("😭 文件或文件夹不存在！")
	}
}
