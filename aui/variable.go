package aui

import (
	g "github.com/AllenDang/giu"
	"grouper/tool"
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

	projects []tool.Project // 项目列表

	isAddProject bool         // 是否添加项目
	oneProject   tool.Project // 需要添加项目

	isFixProject bool         // 是否修改项目
	oldProject   tool.Project // 需要修改的项目
)

// 初始化配置参数
func init() {
	ali = data.UpService.AliyunOss  // 阿里云oss的配置参数
	ten = data.UpService.TencentCos // 腾讯云cos的配置参数
	qin = data.UpService.QiniuOss   // 七牛云oss的配置参数
	projects = data.Projects        // 读取项目列表信息
	ps = initProjectItems()         // 初始化项目列表
}
