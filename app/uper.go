package app

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"grouper/conf"
	"grouper/osser"
	"grouper/tool"
	"log"
	"sync"
	"time"
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

	var wg sync.WaitGroup // 初始化并发池

	var total int // 文件总数
	// 根据不同的配置类型上传
	switch conf.Cfg.UpType {
	case "tencent":
		cos := osser.CosClient() // 腾讯云cos句柄
		p, _ := ants.NewPoolWithFunc(totalPool(len(newPathList)), func(i interface{}) {
			newPath := i.(string)
			couldFile, localFile := name+newPath[len(dirPth):], newPath
			cos.Upload(couldFile, localFile) // 开始上传
			wg.Done()
		}) // 并发任务
		defer p.Release() // 释放并发
		for _, newPath := range newPathList {
			wg.Add(1)
			_ = p.Invoke(newPath) // 执行上传
			if total <= len(newPathList) {
				total++                // 计数
				bar.Play(int64(total)) // 更新进度条
			}
		}
	case "alioss":
		bkt := osser.AliyunGetBucket() // 获取阿里云oss桶
		p, _ := ants.NewPoolWithFunc(totalPool(len(newPathList)), func(i interface{}) {
			newPath := i.(string)
			couldFile, localFile := name+newPath[len(dirPth):], newPath
			bkt.AliyunGoUpload(couldFile, localFile) // 开始上传
			wg.Done()
		}) // 并发任务
		defer p.Release() // 释放并发
		for _, newPath := range newPathList {
			wg.Add(1)
			_ = p.Invoke(newPath) // 执行上传
			if total <= len(newPathList) {
				total++                // 计数
				bar.Play(int64(total)) // 更新进度条
			}
		}
	case "qiniu":
		upt := osser.QiniuGetUpToken() // 获取七牛云上传Token
		p, _ := ants.NewPoolWithFunc(totalPool(len(newPathList)), func(i interface{}) {
			newPath := i.(string)
			couldFile, localFile := name+newPath[len(dirPth):], newPath
			upt.QiniuGoUpload(couldFile, localFile) // 开始上传
			wg.Done()
		}) // 并发任务
		defer p.Release() // 释放并发
		for _, newPath := range newPathList {
			wg.Add(1)
			_ = p.Invoke(newPath) // 执行上传
			if total <= len(newPathList) {
				total++                // 计数
				bar.Play(int64(total)) // 更新进度条
			}
		}
		fmt.Println("💡 淦！自动覆盖已有文件，尚未完成！") // todo 自动覆盖已有文件
	default:
		log.Fatalln("⚠️ 请检查配置文件：up_type 参数不能为空且必须为 tencent alioss qiniu 中的一个。")
	}

	// 执行结束
	wg.Wait()    // 等待并发结束
	bar.Finish() // 结束进度条
	fmt.Printf("🪖 报告长官， %v 个文件上传成功，访问地址为：%v/\n", total, conf.Addr(name))
	fmt.Println("ps: 如果您上传的并非网页文件或图片，可能无法访问哟～")
	// 准备退出
	time.Sleep(3 * time.Second)
	for {
		tool.GoodBye()
	}
}

// 并发数
func totalPool(num int) (total int) {
	switch {
	case num <= 100:
		return 1
	case 101 <= num && num <= 500:
		return 2
	case 501 <= num && num <= 1000:
		return 8
	case 1001 <= num && num <= 5000:
		return 64
	case 5001 <= num && num <= 10000:
		return 128
	case 10001 <= num && num <= 25000:
		return 512
	case 25000 <= num:
		return 1024
	default:
		return 1
	}
}
