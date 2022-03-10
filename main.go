package main

import (
	"flag"
	"fmt"
	"time"
	"upauto/conf"
	"upauto/osser"
	"upauto/tool"
)

// V N 启动参数
var V = flag.String("v", ".", "文件夹的路径，需为绝对路径，默认当前目录")
var N = flag.String("n", "", "name 项目名称，请使用小写字母开头不含特殊符号")

// 初始化
func init() {
	// 解析命令行参数
	flag.Parse()
}

// 程序入口
func main() {
	// 读取启动参数
	dirPth := fmt.Sprintf("%v/%v", *V, *N) // 本地路径
	name := *N                             // 云端文件名

	// 检查命名是否符合规范，文件夹是否存在
	tool.NameStyle(*N, dirPth)

	var total int                     // 文件总数
	bucket := osser.AliyunGetBucket() // 阿里云，获取一个桶子
	// 开始上传，只遍历本地指定的文件夹
	tool.GetFileList(dirPth, func(newPath string) {
		couldFile, localFile := name+newPath[len(dirPth):], newPath
		// 上传程序
		ok := osser.AliyunGoUpload(bucket, couldFile, localFile) // 阿里云
		if ok {
			fmt.Println("上传成功：", localFile, " -> ", couldFile)
			total += 1
		} else {
			fmt.Println("⚠️ 上传失败：", localFile, " -> ", couldFile)
		}
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
