package cli

import (
	"github.com/QiYuTechDev/qiyu-notify-go/api"
	"github.com/QiYuTechDev/qiyu-notify-go/dt"
	"github.com/spf13/cobra"
	"io/ioutil"
)

func init() {
	CliCmd.AddCommand(ddCli)
	ddCli.Flags().StringP("id", "i", "", "钉钉 APP 的 unique_id")
	ddCli.Flags().StringP("msg", "m", "", `推送消息内容
msg 和 file 只需要一个参数即可, 优先使用 msg`)
	ddCli.Flags().StringP("file", "f", "", `推送消息内容
如果没有指定 msg 参数并且指定了 file 参数 则从此 文件中读取`)
}

func ddNotify(uniqueId, msg string) {
	notify := api.ApiDdNotifyUniqueId{}
	_, err := notify.Post(map[string]string{"unique_id": uniqueId}, dt.NotifyArgs{Content: msg})
	if err != nil {
		panic(err)
	}
}

var ddCli = &cobra.Command{
	Use:   "dd",
	Short: "钉钉微信推送",
	Long:  `发送一个通知给钉钉`,
	Run: func(cmd *cobra.Command, args []string) {
		uniqueId, err := cmd.Flags().GetString("id")
		if err != nil {
			panic(err)
		}
		if uniqueId == "" {
			_ = cmd.Help()
			return
		}

		msg, err := cmd.Flags().GetString("msg")
		if err != nil {
			panic(err)
		}

		file, err := cmd.Flags().GetString("file")
		if err != nil {
			panic(err)
		}

		if msg == "" && file != "" {
			content, err := ioutil.ReadFile(file)
			if err != nil {
				panic(err)
			}
			msg = string(content)
		}

		ddNotify(uniqueId, msg)
	},
}
