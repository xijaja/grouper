package aui

import (
	"fmt"
	g "github.com/AllenDang/giu"
	"grouper/app"
	"grouper/conf"
	"sync"
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
func projectList(p *conf.Project) *g.TreeNodeWidget {
	return g.TreeNode(p.Name).Flags(g.TreeNodeFlagsCollapsingHeader).Layout(
		g.Label(fmt.Sprintf("项目名称：%v", p.Name)),
		g.Label(fmt.Sprintf("上传服务：%v", p.UpType)),
		g.Label(fmt.Sprintf("浏览地址：%v", addr(p.Name, p.UpType))).Wrapped(true),
		g.Row(
			// g.Button("拷链").Size(60, 25),
			g.Button("编辑").Size(60, 25).OnClick(func() {
				oldProject = *p // 修改项目参数传递
				if !isFixProject {
					isFixProject = !isFixProject
				} // 如果弹窗未弹出，则使其弹出
			}),
			g.Button("上传").Size(60, 25).OnClick(func() {
				// 如果当前有正在上传的任务
				if isProgressBar {
					g.Msgbox("上传中", "已经有一个链路正在发力...").Buttons(g.MsgboxButtonsOk)
					return
				}
				// 一个新的上传任务
				g.Msgbox("上传项目", "开始打通领域闭环！").Buttons(g.MsgboxButtonsOkCancel).ResultCallback(func(result g.DialogResult) {
					if result {
						// 通过回调获悉，开始上传
						fmt.Println("开始上传……")
						var wg sync.WaitGroup
						wg.Add(1) // 创建一个并发任务
						go func() {
							app.Grouper(*p, upServer(p.UpType), func(n1, n2 int) {
								progressValue = float32(n1) / float32(n2) // 计算进度值
								g.Update()                                // GUI数据更新
							}) // 执行上传
							isProgressBar = false // 上传完成关闭进度条
							progressValue = 0     // 重设进度值
							wg.Done()             // 并发任务完成
						}()
						upLoadProjectName = p.Name // 进度条中显示的项目名
						isProgressBar = true       // 显示上传进度条
					} else {
						fmt.Println("取消上传……")
					}
				})
			}),
		),
	)
}

// 修改项目
func fixOldProject(old *conf.Project) []g.Widget {
	return []g.Widget{
		g.Label(fmt.Sprintf("项目名称：%v", old.Name)), // g.InputText(&old.Name).Size(g.Auto).Flags(g.InputTextFlagsReadOnly),
		g.Dummy(0, 1), // 间隙、空隙
		g.Label("选择上传服务（一定要设置对应的资料哦）"),
		g.Combo("", upType[upTypeSelected], upType, &upTypeSelected).Size(g.Auto).OnChange(func() {
			old.UpType = upType[upTypeSelected]
			fmt.Println(upType[upTypeSelected])
		}),
		g.Label("文件夹路径（文件夹）"),
		g.InputText(&old.LocalFile).Size(g.Auto),
		g.Dummy(0, 3), // 间隙、空隙
		g.Row(
			g.Button("删除").Size(60, 25).OnClick(func() {
				g.Msgbox("删除项目", "快速沉淀，适度倾斜资源").Buttons(g.MsgboxButtonsOkCancel).ResultCallback(func(result g.DialogResult) {
					if result {
						old.DeleteOneProject() // 删除
						isCyclic = true        // 重新读取配置信息
						isFixProject = false   // 关闭项目窗口
					}
				})
			}),
			g.Button("取消").Size(60, 25).OnClick(func() {
				isFixProject = false // 关闭项目窗口
			}),
			g.Button("确定").Size(60, 25).OnClick(func() {
				old.UpdateOneProject() // 修改一个项目
				g.Msgbox("修改完成", "项目重组生态格局").Buttons(g.MsgboxButtonsOk).ResultCallback(func(result g.DialogResult) {
					if result {
						isFixProject = false // 关闭修改项目的窗口
					}
				})
			}),
		),
	}
}

// 添加一个项目 todo 增加对项目名称的判断限制
func addOneProject(one *conf.Project) []g.Widget {
	return []g.Widget{
		g.Label("输入你项目的名字（唯一且非中文，设保存后不可更改）"),
		g.InputText(&one.Name).Size(g.Auto),
		g.Label("选择上传服务（一定要设置对应的资料哦）"),
		g.Combo("", upType[upTypeSelected], upType, &upTypeSelected).Size(g.Auto).OnChange(func() {
			one.UpType = upType[upTypeSelected]
			fmt.Println(upType[upTypeSelected])
		}),
		g.Label("文件夹路径（文件夹）"),
		g.InputText(&one.LocalFile).Size(g.Auto),
		g.Align(g.AlignCenter).To(
			g.Row(
				g.Button("确定").Size(60, 25).OnClick(func() {
					one.AddOneProject() // 添加项目
					fmt.Println("添加项目：", *one)
					g.Msgbox("添加成功", "形成方法论打出一套组合拳").Buttons(g.MsgboxButtonsOk).ResultCallback(func(result g.DialogResult) {
						if result {
							isAddProject = false // 关闭添加项目窗口
						}
					})
				}),
			),
		),
	}
}

// 修改设置
func setUps(any any) []g.Widget {
	switch any.(type) {
	case *conf.AliyunOss: // 修改阿里云oss设置
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
						ali.UpdateAliyunOss() // 更新数据
						fmt.Println("点击保存setUp: ", ali)
						g.Msgbox("保存成功", "已完成对产业链路的赋能升级").Buttons(g.MsgboxButtonsOk).ResultCallback(func(result g.DialogResult) {
							if result {
								isSetUpAli = false // 关闭窗口
							}
						})
					}),
				),
			),
		}
	case *conf.TencentCos: // 修改腾讯云cos设置
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
						ten.UpdateTencentCos() // 更新数据
						fmt.Println("点击保存setUp: ", ten)
						g.Msgbox("保存成功", "已完成对产业链路的赋能升级").Buttons(g.MsgboxButtonsOk).ResultCallback(func(result g.DialogResult) {
							if result {
								isSetUpTen = false // 关闭窗口
							}
						})
					}),
				),
			),
		}
	case *conf.QiniuOss: // 修改七牛云oss设置
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
						qin.UpdateQiniuOss() // 更新数据
						fmt.Println("点击保存setUp: ", qin)
						g.Msgbox("保存成功", "已完成对产业链路的赋能升级").Buttons(g.MsgboxButtonsOk).ResultCallback(func(result g.DialogResult) {
							if result {
								isSetUpQin = false // 关闭窗口
							}
						})
					}),
				),
			),
		}
	default:
		fmt.Println("修改设置面板需要传入被修改参数的指针")
		return nil
	}
}
