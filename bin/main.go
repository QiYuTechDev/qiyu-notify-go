package main

import (
	"fmt"
	"os"
)

func main() {
	if err := cliCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
