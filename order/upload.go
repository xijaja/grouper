package order

import (
	"github.com/spf13/cobra"
	"grouper/app"
	"grouper/conf"
	"grouper/tool"
	"strings"
)

// up 命令
var upCmd = &cobra.Command{
	Use:   "up [flags][name][path]",
	Short: "上传服务",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		qiniu, _ := cmd.Flags().GetBool("qiniu")
		aliyun, _ := cmd.Flags().GetBool("aliyun")
		tencent, _ := cmd.Flags().GetBool("tencent")
		name, _ := cmd.Flags().GetString("name")
		path, _ := cmd.Flags().GetString("path")
		// detail, _ := cmd.Flags().GetBool("detail")
		if !qiniu && !aliyun && !tencent || name == "" {
			_ = cmd.Help()
		} // 如果没有选择任何一个，则提示帮助信息
		// 规范路径
		if name == "" && path == "" {
			cmd.Println("项目名称和路径，均为空")
			_ = cmd.Help()
		}
		if len(path) == 1 && path[len(path)-1:] == "/" {
			// path = path[:len(path)-1] // 去掉最后一个"/"
			path += "/"
		}

		// 规范名称，如果没有指定，则使用路径的最后一个文件夹名
		if name == "" && path != "" {
			arr := strings.Split(path, "/")
			name = arr[len(arr)-2 : len(arr)-1][0]
		}
		tool.NameStyle(name, path) // 检查命名是否符合规范，文件夹是否存在

		// 创建项目配置
		pj := conf.Project{
			Name:      name, // 项目名称
			LocalFile: path, // 本地项目路径
		}
		if qiniu {
			pj.UpType = "七牛云OSS"
			upload := conf.DataInfo.UpService.QiniuOss
			app.CliUper(pj, upload)
			// fmt.Printf("项目信息：%+v\n", pj)     // 打印键值对
			// fmt.Printf("上传信息：%+v\n", upload) // 打印键值对
		} else if aliyun {
			pj.UpType = "阿里云OSS"
			upload := conf.DataInfo.UpService.AliyunOss
			app.CliUper(pj, upload)
		} else if tencent {
			pj.UpType = "腾讯云COS"
			upload := conf.DataInfo.UpService.TencentCos
			app.CliUper(pj, upload)
		} else {
			_ = cmd.Help()
		}
	},
}

// 初始化命令
func init() {
	// 添加到根命令
	rootCmd.AddCommand(upCmd)
	upCmd.Flags().BoolP("help", "h", false, "帮助信息")
	upCmd.Flags().BoolP("qiniu", "q", false, "上传到七牛云")
	upCmd.Flags().BoolP("aliyun", "a", false, "上传到阿里云")
	upCmd.Flags().BoolP("tencent", "t", false, "上传到腾讯云")
	upCmd.Flags().StringP("name", "n", "", "项目名称,应为文件夹名称")
	upCmd.Flags().StringP("path", "p", ".", "本地文件路径")
	upCmd.Flags().BoolP("detail", "d", false, "显示详细上传信息")
}
