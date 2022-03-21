package aui

import (
	"fmt"
	g "github.com/AllenDang/giu"
)

// 遍历项目
func projectItems() (projectsWidget []g.Widget) {
	for _, p := range data.Projects {
		projectsWidget = append(projectsWidget, g.TreeNode(p.Name).Flags(g.TreeNodeFlagsCollapsingHeader).Layout(
			g.Custom(func() {
				if g.IsItemActive() && g.IsMouseClicked(g.MouseButtonLeft) {
					fmt.Println(p.Name, " 的树节点被点击")
				}
			}),
			// g.Selectable("Tree node 1").OnClick(func() {
			// 	fmt.Println("允许被选择")
			// }),
			g.Label(fmt.Sprintf("项目名称：%v", p.Name)),
			g.Label(fmt.Sprintf("上传服务：%v", p.UpType)),
			g.Label(fmt.Sprintf("浏览地址：%v", p.VisitAddr)).Wrapped(true),
			g.Row(
				g.Button("拷贝").Size(60, 25),
				g.Button("修改").Size(60, 25).OnClick(func() {
					isFixProject = !isFixProject
				}), // .OnClick("")
				g.Button("上传").Size(60, 25),
			),
		))
	}
	return projectsWidget
}

// 项目列表
func projectsTabItem() *g.TabItemWidget {
	return g.TabItem("项目列表").Layout(projectItems()...)
}

// 修改项目
func fixProject(isFix bool, projectId int) {
	pj := struct {
		ProjectId int
		Name      string
		UpType    string
		LocalFile string
		VisitAddr string
	}{}
	for _, p := range data.Projects {
		if p.ProjectId == projectId {
			pj.ProjectId = p.ProjectId
			pj.Name = p.Name
			pj.LocalFile = p.LocalFile
			pj.VisitAddr = p.VisitAddr
		}
	}
	if isFix {
		g.Window("修改项目A").IsOpen(&isFix).Flags(g.WindowFlagsNone).Pos(320, 30).Size(300, 150).Layout(
			g.Layout{
				g.Label("输入你项目的名字"),
				g.InputText(&pj.Name).Size(g.Auto),
				g.Label("选择上传服务（一定要设置对应的资料哦）"),
				g.Combo("", upType[upTypeSelected], upType, &upTypeSelected).Size(g.Auto).OnChange(comboChanged),
				g.Label("将文件夹拖放到此窗口（不支持输入）"),
				g.InputText(&pj.LocalFile).Flags(g.InputTextFlagsReadOnly).Size(g.Auto),
			},
			g.Align(g.AlignCenter).To(
				g.Row(
					g.Button("删除").Size(60, 25), // .OnClick(onShowWindow2)
					g.Button("取消").Size(60, 25),
					g.Button("保存").Size(60, 25),
				),
			),
		)
	}
}

// 修改设置
func setUp(isFixSetUp bool) {
	inputSth := data.UpService.AliyunOss
	if isFixSetUp {
		g.Window("阿里云OSS设置").IsOpen(&isFixSetUp).Flags(g.WindowFlagsNone).Pos(90, 110).Size(400, 500).Layout(
			g.Label("endpoint（地域节点地址）"),
			g.InputText(&inputSth.Endpoint).Size(g.Auto),
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

func comboChanged() {
	fmt.Println(upType[upTypeSelected])
}
