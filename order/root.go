package order

import (
	"fmt"
	"github.com/spf13/cobra"
	"grouper/conf"
	"os"
)

// 根命令
var rootCmd = &cobra.Command{
	Use:   "grouper",
	Short: "我的应用程序的简要说明",
	Long:  "cliapp是利用Cobra(Go的CLI库)及Golang开发的命令行应用程序",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
	Version: "1.0.0",
}

// 帮助命令
var rootHelpCmd = &cobra.Command{
	Use:   "help",
	Short: "显示帮助信息",
	Long:  "显示帮助信息",
	Run: func(cmd *cobra.Command, args []string) {
		_ = rootCmd.Help()
	},
}

// Execute 会将所有子命令添加到根命令中，并相应地设置标志。
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// 初始化根命令
func init() {
	// 添加帮助命令
	rootCmd.Flags().BoolP("help", "h", false, "帮助信息")
	rootCmd.SetHelpCommand(rootHelpCmd)
	// 添加版本命令
	rootCmd.Flags().BoolP("version", "v", false, "版本信息")
	rootCmd.SetVersionTemplate(fmt.Sprintf(
		"版本号：%v\n"+
			"开发者：習武（公众号：逆天思维产品汪）\n"+
			"使用说明：xxx.xxx\n"+
			"Github地址：https://github.com/xiwuou/uper\n"+
			"感谢Star 🌟  欢迎Fork 👏\n", conf.Version,
	),
	)
}
