package tpl

func AddCommandTemplate() string {
	return `
package main

import (
	"fmt"

	cmd "github.com/18211167516/go-cmd"
)

// {{ .CmdName }}Cmd represents the {{ .CmdName }} command
var {{ .CmdName }}Cmd = &cmd.Command{
	Use:   "{{ .Use }}",
	Short: "A brief description of your command",
	Long: ` + "`" + `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.` + "`" + `,
	Run: func(Command *cmd.Command, args []string) {
		fmt.Println("{{ .Use }} called")
	},
}

func init() {
	{{ .CmdParent }}.AddCommand({{ .CmdName }}Cmd)
}
`
}
