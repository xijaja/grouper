package order

import (
	"fmt"
	"github.com/spf13/cobra"
	"grouper/conf"
	"grouper/osser"
)

// manager 命令
var managerCmd = &cobra.Command{
	Use:   "manager",
	Short: "对象经理",
	Long: "对象经理：\n" +
		"    文件管理员,目前只支持七牛云,\n" +
		"    查看或删除时每次只处理1000个文件，\n" +
		"    如果文件太多，请分批处理",
	Run: func(cmd *cobra.Command, args []string) {
		prefix, _ := cmd.Flags().GetString("prefix")
		del, _ := cmd.Flags().GetBool("delete")
		if prefix == "" {
			_ = cmd.Help()
		}

		upload := conf.DataInfo.UpService.QiniuOss
		q := osser.QiniuBucketManager(upload)
		var files []string
		if prefix != "" {
			var err error
			files, err = q.GetPrefixFiles(prefix) // 传入文件前缀
			if err != nil {
				cmd.Println("获取文件列表出错:", err)
				return
			}
		}
		if !del {
			for i := 0; i < len(files); i++ {
				fmt.Println(files[i])
			}
		} else {
			fmt.Printf("欲删除文件数：%d\n", len(files))
			err := q.Delete(files)
			if err != nil {
				cmd.Println("删除文件出错:", err)
				return
			}
			fmt.Println("删除成功")
		}
	},
}

func init() {
	rootCmd.AddCommand(managerCmd)
	managerCmd.Flags().BoolP("help", "h", false, "帮助信息")
	managerCmd.Flags().StringP("prefix", "p", "", "查看带有指定前缀的文件")
	managerCmd.Flags().BoolP("delete", "d", false, "删除指定文件")
}
