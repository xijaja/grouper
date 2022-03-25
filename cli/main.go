package main

import (
	"flag"
	"fmt"
	"grouper/conf"
	"grouper/tool"
	"os"
	"strings"
)

// P N 启动参数
var P = flag.String("p", ".", "path 指定上传文件夹的路径，需为绝对路径，默认当前目录")
var N = flag.String("n", "", "name 项目名称，请使用小写字母开头不含特殊符号，默认为文件夹名")

// StartInit 初始化
func main() {
	var printVersion bool // 是否输出版本信息
	flag.BoolVar(&printVersion, "v", false, "显示出版本信息")
	flag.BoolVar(&printVersion, "version", false, "显示出版本信息")
	// 解析命令行参数
	flag.Parse()
	if printVersion {
		fmt.Println("版本号：v1.0-20210311")
		fmt.Println("开发者：習武（公众号：逆天思维产品汪）")
		fmt.Println("使用说明：xxx.xxx")
		fmt.Println("Github地址：https://github.com/xiwuou/uper")
		fmt.Println("感谢Star 🌟  欢迎Fork 👏")
		os.Exit(0) // 退出程序
	}

	// 规范路径
	path := *P
	if path[len(path)-1:] != "/" {
		path = fmt.Sprintf("%v/%v", *P, *N) // 本地路径
	} else {
		path = fmt.Sprintf("%v%v", *P, *N) // 本地路径
	}
	// 规范名称
	name := *N
	if name == "" {
		arr := strings.Split(path, "/")
		name = arr[len(arr)-2 : len(arr)-1][0]
	} // 没有名字则拆分链接最后一个单词
	tool.NameStyle(name, path) // 检查命名是否符合规范，文件夹是否存在
	name += "/"                // 为名字加上斜杠用以命名上传后的文件夹
	// 遍历文件并上传
	pj := conf.Project{
		Name:      name,
		UpType:    "阿里云OSS",
		LocalFile: path,
	}
	upload := conf.DataInfo.UpService.AliyunOss
	fmt.Println(pj)
	fmt.Println(upload)
	// 开始上传
	// app.CliUper(pj, upload)
}
