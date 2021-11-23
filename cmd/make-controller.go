package main

import (
	"go-api/cmd/util"
	"log"
	"os"

	cmd "github.com/18211167516/go-cmd"
)

// MakeControllerCmd represents the MakeController command
var MakeControllerCmd = &cmd.Command{
	Use:   "make:controller",
	Short: "生成 基础版controller",
	Long:  `生成视图、增删改查的controller`,
	Example: `
./cmd.exe make:controller -f test 
会生成到app/controller/testController.go 
model是app/models Test 
service是app/services/TestService
`,
	Run: func(Command *cmd.Command, args []string) {
		file, _ := Command.Flags().GetString("file")
		name, _ := Command.Flags().GetString("name")
		model, _ := Command.Flags().GetString("model")
		service, _ := Command.Flags().GetString("service")
		//
		if serName, err := util.ControllerCreate(file, name, model, service); err != nil {
			log.Printf("[make:controller]-->【%s】生成错误 %v\n", serName, err)
			os.Exit(0)
		} else {
			log.Printf("[make:controller]-->【%s】生成success\n", serName)
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(MakeControllerCmd)
	MakeControllerCmd.Flags().StringP("file", "f", "", "生成的控制器路径加名称,必须指定")
	MakeControllerCmd.Flags().StringP("name", "n", "", "生成控制器的中文注释名")
	MakeControllerCmd.Flags().StringP("model", "m", "", "指定的Model,未指定会按file匹配")
	MakeControllerCmd.Flags().StringP("service", "s", "", "指定的service,未指定会按file匹配")
	MakeControllerCmd.MarkFlagRequired("file")
}
