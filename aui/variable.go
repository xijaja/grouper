package aui

import (
	g "github.com/AllenDang/giu"
	"grouper/tool"
	"time"
)

// 获取配置信息
var (
	data *tool.Data // json配置信息
	ps   []g.Widget // 用以初始化生成项目列表
)

// 修改设置
var (
	isSetUpAli bool            // 是否修改阿里云oss配置
	isSetUpTen bool            // 是否修改腾讯云cos配置
	isSetUpQin bool            // 是否修改七牛云oss配置
	ali        tool.AliyunOss  // 阿里云oss的配置参数
	ten        tool.TencentCos // 腾讯云cos的配置参数
	qin        tool.QiniuOss   // 七牛云oss的配置参数
	// ---------------------------------------------
	projects []tool.Project // 项目列表
	// ---------------------------------------------
	isAddProject bool         // 是否添加项目
	oneProject   tool.Project // 需要添加项目
	// ---------------------------------------------
	isFixProject bool         // 是否修改项目
	oldProject   tool.Project // 需要修改的项目
	// ---------------------------------------------
	upType         []string // 上传服务类型
	upTypeSelected int32    // 上传服务选择
	// ---------------------------------------------
	isProgressBar bool    // 是否显示进度条
	projectName   string  // 正在上传的项目名称
	progressValue float32 // 进度值
)

// 初始化配置参数
func init() {
	data = tool.ReadData()          // json配置信息
	ali = data.UpService.AliyunOss  // 阿里云oss的配置参数
	ten = data.UpService.TencentCos // 腾讯云cos的配置参数
	qin = data.UpService.QiniuOss   // 七牛云oss的配置参数
	projects = data.Projects        // 读取项目列表信息
	ps = initProjectItems()         // 初始化项目列表

	upType = make([]string, 3)
	upType[0] = "阿里云OSS"
	upType[1] = "腾讯云COS"
	upType[2] = "七牛云OSS"
}

// Prg 模拟进度条
func Prg() {
	ticker := time.NewTicker(time.Second * 1)
	progressValue = 0.01
	for {
		if progressValue <= 1 {
			progressValue += 0.05
			g.Update()
			<-ticker.C
		} else {
			break
		}
	}
	isProgressBar = false
}
