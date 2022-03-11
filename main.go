package main

import (
	"flag"
	"fmt"
	"time"
	"upauto/conf"
	"upauto/osser"
	"upauto/tool"
)

// W N 启动参数
var W = flag.String("w", ".", "文件夹的路径，需为绝对路径，默认当前目录")
var N = flag.String("n", "", "name 项目名称，请使用小写字母开头不含特殊符号")

// 初始化
func init() {
	// 版本信息
	var printVersion bool
	flag.BoolVar(&printVersion, "v", false, "显示出版本信息")
	flag.BoolVar(&printVersion, "version", false, "显示出版本信息")
	// 解析命令行参数
	flag.Parse()
	if printVersion {
		conf.Version()
	}
}

// 程序入口
func main() {
	fmt.Printf("example for print version")
	// 读取启动参数
	dirPth := fmt.Sprintf("%v/%v", *W, *N) // 本地路径
	name := *N                             // 云端文件名
	tool.NameStyle(*N, dirPth)             // 检查命名是否符合规范，文件夹是否存在

	var total int            // 文件总数
	cos := osser.CosClient() // 腾讯云cos句柄
	// 开始上传，只遍历本地指定的文件夹
	tool.GetFileList(dirPth, func(newPath string) {
		couldFile, localFile := name+newPath[len(dirPth):], newPath
		fmt.Println("上传成功：", localFile, " -> ", couldFile)
		ok := osser.CosUpload(cos, couldFile, localFile)
		if ok {
			total += 1
		} // 计数
	})

JUDG:
	// 检测是否执行完毕
	agoTotal := total
	time.Sleep(2 * time.Second)
	if agoTotal == total {
		goto OVER // 结束
	} else {
		goto JUDG // 重新检测
	}

OVER:
	// 执行结束
	fmt.Printf("🪖 报告长官！已经上传 %v 个文件，访问地址为：%v/\n", total, conf.Addr(name))
	fmt.Println("ps: 如果您上传的并非网页文件或图片，可能无法访问哟～")
	// 准备退出
	time.Sleep(3 * time.Second)
	for {
		tool.GoodBye()
	}
}
