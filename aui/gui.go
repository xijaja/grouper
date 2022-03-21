package aui

import (
	g "github.com/AllenDang/giu"
	"grouper/tool"
	"os"
)

// 获取配置信息
// var data *tool.Data   // json配置信息
// var isFixProject bool // 是否修改项目
// var isFixSetUp bool
var upType []string
var upTypeSelected int32

// var ali tool.AliyunOss

// 初始化信息
func init() {
	data = tool.ReadData() // json配置信息
	upType = make([]string, 3)
	upType[0] = "阿里云OSS"
	upType[1] = "腾讯云COS"
	upType[2] = "七牛云OSS"
}

// Loop UI界面视图
func Loop() {
	// 顶部导航
	g.MainMenuBar().Layout(
		g.Menu("项目").Layout(
			g.MenuItem("添加").OnClick(func() {
				isAddProject = !isAddProject
			}),
			g.Separator(), // 分割线
			g.MenuItem("退出").OnClick(func() {
				os.Exit(0)
			}),
		),
		g.Menu("设置").Layout(
			// g.Checkbox("复选框", &checked),
			g.MenuItem("阿里云OSS").OnClick(func() {
				isSetUpAli = !isSetUpAli
			}),
			g.MenuItem("腾讯云COS").OnClick(func() {
				isSetUpTen = !isSetUpTen
			}),
			g.MenuItem("七牛云OSS").OnClick(func() {
				isSetUpQin = !isSetUpQin
			}),
		),
		g.Menu("开发者").Layout(
			g.Label("产品：Grouper"),
			g.Label("版本：v1.0.0-beta"),
			g.Label("Github：https://github.com/xiwuou/grouper"),
			g.Separator(), // 分割线
			g.Label("开发者：習武"),
			g.Label("微信号：pm_xiwu（请注明来意）"),
			g.Label("欢迎向我提需求……"),
		),
	).Build()

	// 项目列表（默认）
	g.Window("项目列表").Pos(10, 30).Size(300, 560).Layout(
		g.TabBar().TabItems(
			g.TabItem("项目列表").Layout(
				ps..., // 项目列表
			),
		),
	)

	// 添加项目
	if isAddProject {
		g.Window("添加项目").IsOpen(&isAddProject).Flags(g.WindowFlagsNone).Pos(320, 30).Size(400, 200).Layout(
			addOneProject(&oneProject)...,
		)
	}

	// 修改项目
	if isFixProject {
		g.Window("修改项目").IsOpen(&isFixProject).Flags(g.WindowFlagsNone).Pos(320, 30).Size(400, 200).Layout(
			fixOldProject(&oldProject)...,
		)
	}

	// 修改设置
	if isSetUpAli {
		g.Window("阿里云OSS设置").IsOpen(&isSetUpAli).Flags(g.WindowFlagsNone).Pos(120, 50).Size(400, 500).Layout(
			setUps(&ali)...,
		)
	}
	if isSetUpTen {
		g.Window("腾讯云cos设置").IsOpen(&isSetUpTen).Flags(g.WindowFlagsNone).Pos(130, 60).Size(400, 500).Layout(
			setUps(&ten)...,
		)
	}
	if isSetUpQin {
		g.Window("七牛云oss设置").IsOpen(&isSetUpQin).Flags(g.WindowFlagsNone).Pos(140, 70).Size(400, 500).Layout(
			setUps(&qin)...,
		)
	}
}
