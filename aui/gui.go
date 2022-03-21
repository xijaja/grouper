package aui

import (
	"fmt"
	g "github.com/AllenDang/giu"
	"os"
)

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
		g.PrepareMsgbox(), // 弹窗就绪
		g.TabBar().TabItems(
			g.TabItem("项目列表").Layout(
				ps..., // 项目列表
			),
		),

		// g.Dummy(0, 80), // 间隙、空隙
		// 一个进度条控件 todo 用窗口控件设置
		// g.Align(g.AlignCenter).To(
		// 	g.ProgressBar(0.8),
		// ),
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

	// 上传进度条 todo 同时只能有一个上传任务，即当 isProgressBar 为true时，所有的上传按钮都将被禁用
	if isProgressBar {
		g.Window("正在上传...").IsOpen(&isProgressBar).Flags(g.WindowFlagsNoTitleBar|g.WindowFlagsNoResize|g.WindowFlagsNoCollapse).
			Pos(320, 540).Size(400, 50).Layout(
			g.Align(g.AlignCenter).To(
				g.Label(fmt.Sprintf("【%v】上传中...", projectName)),
				g.ProgressBar(progressValue).Size(g.Auto, 2),
			),
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
