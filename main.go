package main

import (
	"os"

	"wx-cli/internal/wx"
)

func main() {
	os.Exit(wx.Run(os.Args[1:]))
}
