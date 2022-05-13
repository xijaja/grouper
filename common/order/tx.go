package order

import (
	"github.com/spf13/cobra"
	"grouper/common/app"
	"grouper/common/conf"
	"grouper/common/tool"
	"strings"
)

var txCmd = &cobra.Command{
	Use:   "tx",
	Short: "腾讯云 COS 静态云托管",
	Long:  "腾讯云 COS 静态云托管",
	Run: func(cmd *cobra.Command, args []string) {
		list, _ := cmd.Flags().GetBool("list")
		upload, _ := cmd.Flags().GetBool("upload")
		del, _ := cmd.Flags().GetBool("delete")
		name, _ := cmd.Flags().GetString("name")
		path, _ := cmd.Flags().GetString("path")

		if list {
			// 查看项目列表
		} else if upload {
			// 上传项目
			// 如果没有指定项目名称和路径，则提示帮助信息
			if name == "" && path == "" {
				cmd.Println("项目名称和路径，均为空")
				_ = cmd.Help()
			}
			// 如果指定项目名称，但是没有指定路径，则为当前路径
			if len(path) == 1 && path[0] == '.' {
				path = name // 如果没有指定路径，则默认为项目名称
			}
			// 如果指定了路径，但是没有指定项目名称，则使用路径的最后一个文件夹名
			if name == "" && path != "" {
				arr := strings.Split(path, "/")
				name = arr[len(arr)-2 : len(arr)-1][0]
			}
			tool.NameStyle(name, path) // 检查命名是否符合规范，文件夹是否存在
			// 开始上传
			cmd.Println("正在扫描本地文件，准备上传到阿里云OSS...")
			app.CliUper(conf.Project{
				Name:      name,     // 项目名称
				LocalFile: path,     // 本地项目路径
				UpType:    "腾讯云COS", // 上传服务类型
			}, conf.DataInfo.UpService.TencentCos)
		} else if del {
			// 删除项目
		} else {
			// 无效的命令
			_ = cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(txCmd)
	txCmd.Flags().BoolP("help", "h", false, "帮助信息")
	txCmd.Flags().BoolP("list", "l", false, "查看")
	txCmd.Flags().BoolP("delete", "d", false, "删除")
	txCmd.Flags().BoolP("upload", "u", false, "上传")
	txCmd.Flags().StringP("name", "n", "", "项目名称,应为文件夹名称")
	txCmd.Flags().StringP("path", "p", ".", "本地文件路径")
}
