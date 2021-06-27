package cli

import (
	"fmt"
	"github.com/QiYuTechDev/qiyu-notify-go/api"
	"github.com/QiYuTechDev/qiyu-notify-go/dt"
	"github.com/spf13/cobra"
	"io/ioutil"
)

const (
	typTxt = "txt"
	typMd  = "md"
)

func init() {
	CliCmd.AddCommand(wxCli)
	wxCli.Flags().StringP("id", "i", "", "企业微信的 unique_id")
	wxCli.Flags().StringP("type", "t", "txt", `推送消息类型
当前支持:
* txt 纯文本
* md: 企业微信 markdown 类型`)
	wxCli.Flags().StringP("msg", "m", "", `推送消息内容
msg 和 file 只需要一个参数即可, 优先使用 msg`)
	wxCli.Flags().StringP("file", "f", "", `推送消息内容
如果没有指定 msg 参数并且指定了 file 参数 则从此 文件中读取`)
}

func wxNotify(uniqueId, typ, msg string) {
	if typ == typTxt {
		notify := api.ApiWxNotifyUniqueId{}
		_, err := notify.Post(map[string]string{"unique_id": uniqueId}, dt.NotifyArgs{Content: msg})
		if err != nil {
			panic(err)
		}
	} else {
		notify := api.ApiWxMdUniqueId{}
		_, err := notify.Post(map[string]string{"unique_id": uniqueId}, dt.NotifyArgs{Content: msg})
		if err != nil {
			panic(err)
		}
	}
}

var wxCli = &cobra.Command{
	Use:   "wx",
	Short: "企业微信推送",
	Long:  `发送一个通知给企业微信`,
	Run: func(cmd *cobra.Command, args []string) {
		uniqueId, err := cmd.Flags().GetString("id")
		if err != nil {
			panic(err)
		}
		if uniqueId == "" {
			_ = cmd.Help()
			return
		}

		typ, err := cmd.Flags().GetString("type")
		if err != nil {
			panic(err)
		}
		if typ != typTxt && typ != typMd {
			panic(fmt.Sprintf("type: %s 是无效值!", typ))
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

		wxNotify(uniqueId, typ, msg)
	},
}
