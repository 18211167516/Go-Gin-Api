package main

import (
	"go-api/cmd/util"
	"log"
	"os"

	cmd "github.com/18211167516/go-cmd"
)

// MakeViewCmd represents the MakeView command
var MakeViewCmd = &cmd.Command{
	Use:   "make:view",
	Short: "生成视图层默认页面",
	Long:  `支持生成视图`,
	Run: func(Command *cmd.Command, args []string) {
		path, _ := Command.Flags().GetString("path")
		desc, _ := Command.Flags().GetString("desc")
		if err := util.ViewCreate(path, desc); err != nil {
			log.Printf("[make:view]-->【%s】生成错误 %v\n", path, err)
			os.Exit(0)
		} else {

			log.Printf("[make:view]-->【%s】生成success\n", path)
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(MakeViewCmd)
	MakeViewCmd.Flags().StringP("path", "p", "", "生成的目录例如test/test_list 会在templates/test 生成test_list.html")
	MakeViewCmd.Flags().StringP("desc", "d", "", "文件描述")
	MakeViewCmd.MarkFlagRequired("path")
	MakeViewCmd.MarkFlagRequired("desc")
}
