package app

import (
	"fmt"
	"log"
	"time"
	"upauto/conf"
	"upauto/osser"
	"upauto/tool"
)

func Uper() {
	// 读取启动参数
	dirPth, name := conf.CheckStartup()
	// 遍历本地指定的文件夹  文件路径列表
	newPathList := tool.GetFileList(dirPth)
	// 上传过程的进度条
	var bar tool.Bar
	bar.NewOption(0, int64(len(newPathList)))
	fmt.Println("扫描完成，开始上传：")

	var total int // 文件总数
	// 根据不同的配置类型上传
	switch conf.Cfg.UpType {
	case "tencent":
		cos := osser.CosClient() // 腾讯云cos句柄
		for i, newPath := range newPathList {
			couldFile, localFile := name+newPath[len(dirPth):], newPath
			cos.Upload(couldFile, localFile) // 开始上传
			if total <= len(newPathList) {
				bar.Play(int64(i + 1))
			}
			total++
		}
	case "alioss":
		bkt := osser.AliyunGetBucket() // 获取阿里云oss桶
		for i, newPath := range newPathList {
			couldFile, localFile := name+newPath[len(dirPth):], newPath
			bkt.AliyunGoUpload(couldFile, localFile)
			if total <= len(newPathList) {
				bar.Play(int64(i + 1))
			}
			total++
		}
	case "qiniu":
		upt := osser.QiniuGetUpToken() // 获取七牛云上传Token
		for i, newPath := range newPathList {
			couldFile, localFile := name+newPath[len(dirPth):], newPath
			upt.QiniuGoUpload(couldFile, localFile)
			if total <= len(newPathList) {
				bar.Play(int64(i + 1))
			}
			total++
		}
		fmt.Println("💡 淦！自动覆盖已有文件，尚未完成！") // todo 自动覆盖已有文件
	default:
		log.Fatalln("⚠️ 请检查配置文件：up_type 参数不能为空且必须为 tencent alioss qiniu 中的一个。")
	}

	// 结束进度条
	bar.Finish()

	// 执行结束
	fmt.Printf("🪖 报告长官， %v 个文件上传成功，访问地址为：%v/\n", total, conf.Addr(name))
	fmt.Println("ps: 如果您上传的并非网页文件或图片，可能无法访问哟～")
	// 准备退出
	time.Sleep(3 * time.Second)
	for {
		tool.GoodBye()
	}
}
