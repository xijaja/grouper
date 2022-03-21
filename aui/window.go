package aui

import (
	"fmt"
	g "github.com/AllenDang/giu"
	"grouper/tool"
)

// 遍历项目（初始化时遍历）
func initProjectItems() (projectsWidget []g.Widget) {
	num := len(projects)
	for i := 0; i < num; i++ {
		pro := projectList(&projects[i])
		projectsWidget = append(projectsWidget, pro)
	}
	return projectsWidget
}

// 项目列表
func projectList(p *tool.Project) *g.TreeNodeWidget {
	return g.TreeNode(p.Name).Flags(g.TreeNodeFlagsCollapsingHeader).Layout(
		g.Label(fmt.Sprintf("项目名称：%v", p.Name)),
		g.Label(fmt.Sprintf("上传服务：%v", p.UpType)),
		g.Label(fmt.Sprintf("浏览地址：%v", p.VisitAddr)).Wrapped(true),
		g.Row(
			g.Button("拷链").Size(60, 25), // todo 拷贝链接
			g.Button("修改").Size(60, 25).OnClick(func() {
				oldProject = *p // 参数传递
				if !isFixProject {
					isFixProject = !isFixProject
				} // 如果弹窗未弹出，则使其弹出
			}),
			g.Button("上传").Size(60, 25).OnClick(func() {
				// 如果当前有正在上传的任务
				if isProgressBar {
					g.Msgbox("心急吃不了热豆腐", "已经有一个任务正在执行了...").Buttons(g.MsgboxButtonsOk)
					return
				}
				// 一个新的上传任务
				g.Msgbox("上传项目", "准备好开始上传了吗？").Buttons(g.MsgboxButtonsOkCancel).
					ResultCallback(func(result g.DialogResult) {
						if result {
							// 通过回调获悉，开始上传
							fmt.Println("开始上传……")
							projectName = p.Name
							isProgressBar = true
						} else {
							fmt.Println("取消上传……")
						}
					})
				return
			}),
		),
	)
}

// 修改项目 todo 保存写入并弹窗
func fixOldProject(old *tool.Project) []g.Widget {
	return []g.Widget{
		g.Label("输入你项目的名字（注意，唯一且非中文）"), // todo 项目名应唯一且非中文
		g.InputText(&old.Name).Size(g.Auto),
		g.Label("选择上传服务（一定要设置对应的资料哦）"),
		g.Combo("", upType[upTypeSelected], upType, &upTypeSelected).Size(g.Auto).OnChange(func() {
			fmt.Println(upType[upTypeSelected])
		}),
		g.Label("将文件夹拖放到此窗口（不支持输入）"),
		g.InputText(&old.LocalFile).Flags(g.InputTextFlagsReadOnly).Size(g.Auto),
		g.Align(g.AlignCenter).To(
			g.Row(
				g.Button("确定").Size(60, 25).OnClick(func() {
					fmt.Println("点击确定")
					fmt.Println(*old)
				}),
			),
		),
	}
}

// 添加一个项目 todo 保存写入并弹窗
func addOneProject(one *tool.Project) []g.Widget {
	return []g.Widget{
		g.Label("输入你项目的名字（注意，这不应该是中文）"), // todo 项目名应唯一且非中文
		g.InputText(&one.Name).Size(g.Auto),
		g.Label("选择上传服务（一定要设置对应的资料哦）"),
		g.Combo("", upType[upTypeSelected], upType, &upTypeSelected).Size(g.Auto).OnChange(func() {
			fmt.Println(upType[upTypeSelected])
		}),
		g.Label("将文件夹拖放到此窗口（不支持输入）"),
		g.InputText(&one.LocalFile).Flags(g.InputTextFlagsReadOnly).Size(g.Auto),
		g.Align(g.AlignCenter).To(
			g.Row(
				g.Button("确定").Size(60, 25).OnClick(func() {
					fmt.Println("点击确定")
					fmt.Println(*one)
				}),
			),
		),
	}
}

// 修改设置 todo 保存写入并弹窗
func setUps(any any) []g.Widget {
	switch any.(type) {
	case *tool.AliyunOss: // 修改阿里云oss设置
		return []g.Widget{
			g.Label("endpoint（地域节点地址）"),
			g.InputText(&ali.Endpoint).Size(g.Auto),
			g.Label("key_id（oss的key）"),
			g.InputText(&ali.KeyID).Size(g.Auto),
			g.Label("key_secret（oss的secret）"),
			g.InputText(&ali.KeySecret).Size(g.Auto),
			g.Label("bucket_name（储存桶的名字）"),
			g.InputText(&ali.BucketName).Size(g.Auto),
			g.Label("visit_addr（绑定的域名或查看地址）"),
			g.InputText(&ali.Domain).Size(g.Auto),
			g.Align(g.AlignCenter).To(
				g.Row(
					g.Button("保存").Size(60, 25).OnClick(func() {
						fmt.Println("点击保存setUp")
						fmt.Println(ali)
					}),
				),
			),
		}
	case *tool.TencentCos: // 修改腾讯云cos设置
		return []g.Widget{
			g.Label("bucket_name（桶名）"),
			g.InputText(&ten.BucketName).Size(g.Auto),
			g.Label("cos_region（区域）"),
			g.InputText(&ten.CosRegion).Size(g.Auto),
			g.Label("secret_id（id）"),
			g.InputText(&ten.SecretID).Size(g.Auto),
			g.Label("secret_key（key）"),
			g.InputText(&ten.SecretKey).Size(g.Auto),
			g.Align(g.AlignCenter).To(
				g.Row(
					g.Button("保存").Size(60, 25).OnClick(func() {
						fmt.Println("点击保存setUp")
						fmt.Println(ten)
					}),
				),
			),
		}
	case *tool.QiniuOss: // 修改七牛云oss设置
		return []g.Widget{
			g.Label("bucket_name（桶名）"),
			g.InputText(&qin.BucketName).Size(g.Auto),
			g.Label("access_key（ak）"),
			g.InputText(&qin.AccessKey).Size(g.Auto),
			g.Label("secret_key（sk）"),
			g.InputText(&qin.SecretKey).Size(g.Auto),
			g.Align(g.AlignCenter).To(
				g.Row(
					g.Button("保存").Size(60, 25).OnClick(func() {
						fmt.Println("点击保存setUp")
						fmt.Println(qin)
					}),
				),
			),
		}
	default:
		fmt.Println("修改设置面板需要传入被修改参数的指针")
		return nil
	}
}

func comboChanged() {
	fmt.Println(upType[upTypeSelected])
}
