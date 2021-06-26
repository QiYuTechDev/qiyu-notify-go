package cli

import (
	"github.com/spf13/cobra"
)

var CliCmd = &cobra.Command{
	Short: "notify 奇遇推送命令行工具",
	Long: `notify 是一个在任意地方给您发送通知的命令行工具

当前仅支持奇遇推送官方的服务器[未来会支持自定义服务器]
`,
}
