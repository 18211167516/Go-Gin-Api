package main

import (
	cmd "github.com/18211167516/go-cmd"
	"github.com/gookit/color"
)

// versionCmd represents the version command
var versionCmd = &cmd.Command{
	Use:   "version",
	Short: "版本信息",
	Long:  `版本的长信息`,
	Run: func(Command *cmd.Command, args []string) {
		color.Info.Println("v0.0.1")
	},
}

func init() {
	cmd.RootCmd.AddCommand(versionCmd)
}
