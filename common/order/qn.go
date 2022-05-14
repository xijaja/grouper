package order

import (
	"fmt"
	"github.com/spf13/cobra"
	"grouper/common/app"
	"grouper/common/conf"
	"grouper/common/osser"
	"grouper/common/tool"
	"strings"
)

var qnCmd = &cobra.Command{
	Use:   "qn",
	Short: "七牛云 OSS 静态云托管",
	Long:  "七牛云 OSS 静态云托管",
	Run: func(cmd *cobra.Command, args []string) {
		list, _ := cmd.Flags().GetBool("list")
		upload, _ := cmd.Flags().GetBool("upload")
		del, _ := cmd.Flags().GetBool("delete")
		name, _ := cmd.Flags().GetString("name")
		path, _ := cmd.Flags().GetString("path")
		prefix, _ := cmd.Flags().GetString("prefix")

		if list {
			// 查看项目列表
			if prefix == "" {
				_ = cmd.Help() // 没有指定前缀，则提示帮助信息
			}
			qnOss := conf.DataInfo.UpService.QiniuOss // 获取七牛云配置
			q := osser.QiniuBucketManager(qnOss)      // 初始化七牛云 OSS
			var files []string                        // 文件列表
			files, err := q.GetPrefixFiles(prefix)    // 传入文件前缀
			if err != nil {
				cmd.Println("获取文件列表出错:", err)
				return
			}
			for i := 0; i < len(files); i++ {
				cmd.Println(files[i])
			}
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
				UpType:    "七牛云OSS", // 上传服务类型
			}, conf.DataInfo.UpService.QiniuOss)
		} else if del {
			// 删除项目(七牛云指定前缀)
			if prefix == "" {
				_ = cmd.Help() // 没有指定前缀，则提示帮助信息
			}
			qnOss := conf.DataInfo.UpService.QiniuOss // 获取七牛云配置
			q := osser.QiniuBucketManager(qnOss)      // 初始化七牛云 OSS
			var files []string                        // 文件列表
			files, err := q.GetPrefixFiles(prefix)    // 传入文件前缀
			if err != nil {
				cmd.Println("获取文件列表出错:", err)
				return
			}
			fmt.Printf("欲删除文件数：%d\n", len(files))
			err = q.Delete(files)
			if err != nil {
				cmd.Println("删除文件出错:", err)
				return
			}
			fmt.Println("删除成功")
		} else {
			// 无效的命令
			_ = cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(qnCmd)
	qnCmd.Flags().BoolP("help", "h", false, "帮助信息")
	qnCmd.Flags().BoolP("list", "l", false, "查看")
	qnCmd.Flags().BoolP("delete", "d", false, "删除")
	qnCmd.Flags().BoolP("upload", "u", false, "上传")
	qnCmd.Flags().StringP("prefix", "f", "", "查看带有指定前缀的文件")
	qnCmd.Flags().StringP("name", "n", "", "项目名称,应为文件夹名称")
	qnCmd.Flags().StringP("path", "p", ".", "本地文件路径")
}
