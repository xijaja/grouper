package aui

import (
	"fmt"
	g "github.com/AllenDang/giu"
	"grouper/tool"
	"os"
)

// 获取配置信息
var data *tool.Data   // json配置信息
var isFixProject bool // 是否修改项目
var isFixSetUp bool
var upType []string
var upTypeSelected int32

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
				// showPD = !showPD
			}),
			g.Separator(), // 分割线
			g.MenuItem("退出").OnClick(func() {
				os.Exit(0)
			}),
		),
		g.Menu("设置").Layout(
			// g.Checkbox("复选框", &checked),
			g.MenuItem("阿里云OSS").OnClick(func() {
				isFixSetUp = !isFixSetUp
			}),
			g.MenuItem("腾讯云COS"),
			g.MenuItem("七牛云OSS"),
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
	g.Window("项目列表").Pos(10, 30).IsOpen(nil).Size(300, 560).Layout(
		g.TabBar().TabItems(
			projectsTabItem(), // 项目列表
		),
	)

	// 修改项目
	fixProject(isFixProject, 2)

	// 修改
	inputSth := data.UpService.AliyunOss
	if isFixSetUp {
		g.Window("阿里云OSS设置").IsOpen(&isFixSetUp).Flags(g.WindowFlagsNone).Pos(90, 110).Size(400, 500).Layout(
			g.Label("endpoint（地域节点地址）"),
			g.InputText(&inputSth.Endpoint).Size(g.Auto),
			g.InputTextMultiline(&inputSth.Endpoint),
			g.Label("key_id（oss的key）"),
			g.InputText(&inputSth.KeyID).Size(g.Auto),
			g.Label("key_secret（oss的secret）"),
			g.InputText(&inputSth.KeySecret).Size(g.Auto),
			g.Label("bucket_name（储存桶的名字）"),
			g.InputText(&inputSth.BucketName).Size(g.Auto),
			g.Label("visit_addr（绑定的域名或查看地址）"),
			g.InputText(&inputSth.Domain).Size(g.Auto),
			g.Align(g.AlignCenter).To(
				g.Row(
					g.Button("保存").Size(60, 25).OnClick(func() {
						fmt.Println("点击保存setUp")
						fmt.Println(inputSth)
					}),
				),
			),
		)
	}
}
