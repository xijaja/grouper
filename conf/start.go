package conf

import (
	"flag"
	"fmt"
	"os"
	"upauto/tool"
)

// ---------------------------------------------
// 启动信息
// ---------------------------------------------

// P N 启动参数
var P = flag.String("p", ".", "path 文件夹的路径，需为绝对路径，默认当前目录")
var N = flag.String("n", "", "name 项目名称，请使用小写字母开头不含特殊符号")

// 初始化
func init() {
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
}

// CheckStartup 检查启动参数
func CheckStartup() (path, name string) {
	p := *P
	if p[len(p)-1:] != "/" {
		path = fmt.Sprintf("%v/%v", *P, *N) // 本地路径
	} else {
		path = fmt.Sprintf("%v%v", *P, *N) // 本地路径
	}
	name = *N                // 文件名，云端&云端
	tool.NameStyle(*N, path) // 检查命名是否符合规范，文件夹是否存在
	return path, name
}
