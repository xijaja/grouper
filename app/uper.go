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

// CliUper 是CLI版
// ---------------------------------------------
func CliUper(project conf.Project, upServer any) {
	// 声明进度条
	var bar tool.Bar
	// 执行上传
	var ts int
	num, addr := Grouper(project, upServer, func(n1, n2 int) {
		if ts == 0 {
			bar.NewOption(0, int64(n2)) // 创建进度条
			bar.Play(int64(n1))         // 进度值
		} else {
			bar.Play(int64(n1)) // 进度值
		}
		ts++ // 更新被调用次数
	})
	bar.Finish() // 结束进度条
	// 执行结束
	fmt.Printf("🪖 报告长官， %v 个文件上传成功，访问地址为：%v/\n", num, addr)
	fmt.Println("ps: 如果您上传的并非网页文件或图片，可能无法访问哟～")
	// 结束退出
	if !tool.IsOsWindows() {
		return
	} else {
		time.Sleep(3 * time.Second)
		for {
			tool.GoodBye()
		}
	}
}

// Grouper 是GUI版
// ---------------------------------------------
func Grouper(project conf.Project, upServer any, f func(n1, n2 int)) (num int, addr string) {
	dirPth, name := project.LocalFile, project.Name
	newPathList := tool.GetFileList(dirPth) // 遍历本地指定的文件夹，文件路径列表
	fmt.Println("扫描完成，开始上传：")

	var wg sync.WaitGroup // 初始化并发池
	var total int         // 已上传的文件总数
	var domain string     // 查看地址的域名
	switch project.UpType {
	case "阿里云OSS":
		ali := upServer.(conf.AliyunOss)
		bkt := osser.AliyunGetBucket(ali) // 获取阿里云oss桶
		p, _ := ants.NewPoolWithFunc(totalPool(len(newPathList)), func(i interface{}) {
			newPath := i.(string)
			couldFile, localFile := name+"/"+newPath[len(dirPth):], newPath
			bkt.AliyunGoUpload(couldFile, localFile) // 开始上传
			wg.Done()
		}) // 并发任务
		defer p.Release() // 释放并发
		for _, newPath := range newPathList {
			wg.Add(1)
			_ = p.Invoke(newPath) // 执行上传
			if total <= len(newPathList) {
				total++                    // 计数
				f(total, len(newPathList)) // 进度回调
			}
		}
		domain = ali.Domain
	case "腾讯云COS":
		tx := upServer.(conf.TencentCos)
		cos := osser.CosClient(tx) // 腾讯云cos句柄
		p, _ := ants.NewPoolWithFunc(totalPool(len(newPathList)), func(i interface{}) {
			newPath := i.(string)
			couldFile, localFile := name+"/"+newPath[len(dirPth):], newPath
			cos.Upload(couldFile, localFile) // 开始上传
			wg.Done()
		}) // 并发任务
		defer p.Release() // 释放并发
		for _, newPath := range newPathList {
			wg.Add(1)
			_ = p.Invoke(newPath) // 执行上传
			if total <= len(newPathList) {
				total++                    // 计数
				f(total, len(newPathList)) // 进度回调
			}
		}
		domain = tx.Domain
	case "七牛云OSS":
		qin := upServer.(conf.QiniuOss)
		upt := osser.QiniuGetUpToken(qin) // 获取七牛云上传Token
		p, _ := ants.NewPoolWithFunc(totalPool(len(newPathList)), func(i interface{}) {
			newPath := i.(string)
			couldFile, localFile := name+"/"+newPath[len(dirPth):], newPath
			upt.QiniuGoUpload(couldFile, localFile) // 开始上传
			wg.Done()
		}) // 并发任务
		defer p.Release() // 释放并发
		for _, newPath := range newPathList {
			wg.Add(1)
			_ = p.Invoke(newPath) // 执行上传
			if total <= len(newPathList) {
				total++                    // 计数
				f(total, len(newPathList)) // 进度回调
			}
		}
		domain = qin.Domain
		fmt.Println("💡 淦！自动覆盖已有文件，尚未完成！") // todo 自动覆盖已有文件
	default:
		log.Fatalln("⚠️ 请检查配置文件")
	}
	// 执行结束
	wg.Wait() // 等待并发结束
	fmt.Println("上传完成！")
	viewAddr := view(name, domain)
	return total, viewAddr
}

// ---------------------------------------------

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

// 查看地址
func view(name, domain string) (addr string) {
	if domain == "" {
		return name
	}
	if domain[len(domain)-1:] != "/" {
		return domain + "/" + name
	} else {
		return domain + name
	}
}
