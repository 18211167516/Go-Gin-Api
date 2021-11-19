package main

import (
	"fmt"
	"go-api/tool"
	"log"
	"os"
	"strings"
	"text/template"

	cmd "github.com/18211167516/go-cmd"
)

type command struct {
	Use       string
	FileName  string
	CmdName   string
	CmdParent string
}

// versionCmd represents the version command
var makeCommandCmd = &cmd.Command{
	Use:   "make:command",
	Short: "生成命令",
	Long:  `生成自定义命令`,
	Run: func(Command *cmd.Command, args []string) {
		c, _ := Command.Flags().GetString("name")
		cmdName := tool.CamelCase(c, "-")
		use := strings.Replace(c, "-", ":", -1)
		st := &command{
			Use:       use,
			FileName:  c,
			CmdName:   cmdName,
			CmdParent: "cmd.RootCmd",
		}

		if err := st.Create(); err != nil {
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

//创建命令文件
func (c *command) Create() error {
	cmdFilePath := fmt.Sprintf("%s.go", c.FileName)
	cmdFile, err := os.Create(cmdFilePath)
	if err != nil {
		return err
	}
	defer cmdFile.Close()

	commandTemplate := template.Must(template.New("sub").Parse(AddCommandTemplate()))
	err = commandTemplate.Execute(cmdFile, c)
	if err != nil {
		return err
	}
	return nil
}

//模板
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
