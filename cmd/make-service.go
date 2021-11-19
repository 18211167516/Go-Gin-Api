package main

import (
	"go-api/cmd/util"
	"log"
	"os"

	cmd "github.com/18211167516/go-cmd"
)

// MakeServiceCmd represents the MakeService command
var MakeServiceCmd = &cmd.Command{
	Use:   "make:service",
	Short: "快捷生成service",
	Long:  `支持快捷生成标准化serveice`,
	Example: `
根据app/models/test Test
生成默认文件 TestService.go
./cmd.exe make:service -m=test/Test 
生成自定义文件 testSerive.go
./cmd.exe make:service -m=test/Test -f=testService
生成到自定义目录
./cmd.exe make:service -m=test/Test -p=test
`,
	Run: func(Command *cmd.Command, args []string) {
		model, _ := Command.Flags().GetString("model")
		file, _ := Command.Flags().GetString("file")
		path, _ := Command.Flags().GetString("path")
		//
		if serName, err := util.ServiceCreate(model, file, path); err != nil {
			log.Printf("[make:service]-->【%s】生成错误 %v\n", serName, err)
			os.Exit(0)
		} else {

			log.Printf("[make:service]-->【%s】生成success\n", serName)
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(MakeServiceCmd)
	MakeServiceCmd.Flags().StringP("model", "m", "", "指定的Model,必须指定")
	MakeServiceCmd.Flags().StringP("file", "f", "", "生成的文件名，如果未指定则会通过model")
	MakeServiceCmd.Flags().StringP("path", "p", "", "生成的目录")
	MakeServiceCmd.MarkFlagRequired("model")
}
