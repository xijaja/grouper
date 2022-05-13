package aui

import (
	g "github.com/AllenDang/giu"
	"grouper/common/conf"
)

// 获取配置信息
var (
	data     *conf.Data // json配置信息
	ps       []g.Widget // 用以初始化生成项目列表
	isCyclic bool       // 是否重新读取数据
	// ---------------------------------------------
	isSetUpAli bool            // 是否修改阿里云oss配置
	isSetUpTen bool            // 是否修改腾讯云cos配置
	isSetUpQin bool            // 是否修改七牛云oss配置
	ali        conf.AliyunOss  // 阿里云oss的配置参数
	ten        conf.TencentCos // 腾讯云cos的配置参数
	qin        conf.QiniuOss   // 七牛云oss的配置参数
	// ---------------------------------------------
	projects []conf.Project // 项目列表
	// ---------------------------------------------
	isAddProject bool         // 是否添加项目
	oneProject   conf.Project // 需要添加项目
	// ---------------------------------------------
	isFixProject bool         // 是否修改项目
	oldProject   conf.Project // 需要修改的项目
	// ---------------------------------------------
	upType         []string // 上传服务类型
	upTypeSelected int32    // 上传服务选择
	// ---------------------------------------------
	isProgressBar     bool    // 是否显示进度条
	upLoadProjectName string  // 正在上传的项目名称
	progressValue     float32 // 进度值
)

// 初始化配置参数
func init() {
	data = conf.DataInfo            // json配置信息
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

// 循环更新读取数据
func cyclicUpdate() {
	data = conf.RereadDataInfo()    // json配置信息（再次读取时数据在内存中，无需从配置文件获取）
	ali = data.UpService.AliyunOss  // 阿里云oss的配置参数
	ten = data.UpService.TencentCos // 腾讯云cos的配置参数
	qin = data.UpService.QiniuOss   // 七牛云oss的配置参数
	projects = data.Projects        // 读取项目列表信息
	ps = initProjectItems()         // 初始化项目列表
	g.Update()
}

// addr 访问地址
func addr(name string, upType string) (addr string) {
	switch upType {
	case "阿里云OSS":
		addr = ali.Domain
	case "腾讯云COS":
		addr = ten.Domain
	case "七牛云OSS":
		addr = qin.Domain
	}
	if addr == "" {
		return name
	}
	if addr[len(addr)-1:] != "/" {
		return addr + "/" + name
	} else {
		return addr + name
	}
}

// 上传服务，绑定对应的配置
func upServer(upType string) any {
	switch upType {
	case "阿里云OSS":
		return ali
	case "腾讯云COS":
		return ten
	case "七牛云OSS":
		return qin
	default:
		return ali
	}
}

// 点击上传时，判断对应配置是否完整
func checkupServerConfig(upType string) (isComplete bool) {
	switch upType {
	case "阿里云OSS":
		if ali.KeyID == "" || ali.KeySecret == "" || ali.Endpoint == "" || ali.BucketName == "" {
			return false
		}
		return true
	case "腾讯云COS":
		if ten.BucketName == "" || ten.SecretID == "" || ten.SecretKey == "" || ten.CosRegion == "" {
			return false
		}
		return true
	case "七牛云OSS":
		if qin.BucketName == "" || qin.AccessKey == "" || qin.SecretKey == "" {
			return false
		}
		return true
	default:
		return false
	}
}

// 更新文本信息
// wnd.SetDropCallback(onDrop)
// func onDrop(names []string) {
// 	var sb strings.Builder
// 	for _, n := range names {
// 		path := fmt.Sprintf("%s", n)
// 		sb.WriteString(path)
// 		fmt.Println(path)
// 	}
// 	dropInFiles = sb.String()
// 	g.Update()
// }
