package main

import (
	"log"

	cmd "github.com/18211167516/go-cmd"
)

// versionCmd represents the version command
var versionCmd = &cmd.Command{
	Use:   "version",
	Short: "版本信息",
	Long:  `版本的长信息`,
	Run: func(Command *cmd.Command, args []string) {
		log.Println("v0.0.2")
	},
}

func init() {
	cmd.RootCmd.AddCommand(versionCmd)
}
