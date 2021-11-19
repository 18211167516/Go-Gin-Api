package main

import (
	"go-api/cmd/util"
	"go-api/tool"
	"log"
	"os"
	"strings"

	cmd "github.com/18211167516/go-cmd"
)

// versionCmd represents the version command
var makeCommandCmd = &cmd.Command{
	Use:   "make:command",
	Short: "生成命令",
	Long:  `生成自定义命令`,
	Example: `
生成make-service命令：同时会将-转成:
./cmd.exe make:command -n=make-service
`,
	Run: func(Command *cmd.Command, args []string) {
		c, _ := Command.Flags().GetString("name")
		cmdName := tool.CamelCase(c, "-")
		use := strings.Replace(c, "-", ":", -1)

		if err := util.CommandCreate(use, c, cmdName, "cmd.RootCmd"); err != nil {
			log.Printf("[make:command]-->命令【%sCmd】生成错误 %v\n", use, err)
			os.Exit(0)
		}
		log.Printf("[make:command]-->命令【%sCmd】生成success\n", use)
	},
}

func init() {
	cmd.RootCmd.AddCommand(makeCommandCmd)
	makeCommandCmd.Flags().StringP("name", "n", "", "命令名称")
	makeCommandCmd.MarkFlagRequired("name")
}
