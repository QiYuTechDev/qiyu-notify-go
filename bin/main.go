package main

import (
	"fmt"
	"github.com/QiYuTechDev/qiyu-notify-go/cli"
	"os"
)

func main() {
	if err := cli.CliCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
