package cli

import (
	"github.com/QiYuTechDev/qiyu-notify-go/api"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	CliCmd.AddCommand(pingCli)
}

var pingCli = &cobra.Command{
	Use:   "ping",
	Short: "测试服务器是否正常",
	Long:  `测试推送服务器是否可以正常工作`,
	Run: func(cmd *cobra.Command, args []string) {
		ping := api.ApiPing{}
		data, err := ping.Get()
		if err != nil {
			panic(err)
		}

		s := string(data.([]byte))
		if s == "pong" {
			println("服务器运行正常")
		} else {
			println("服务器异常: ", s)
			os.Exit(1)
		}
	},
}
