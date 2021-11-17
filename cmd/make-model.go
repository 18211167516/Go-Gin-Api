package main

import (
	"go-api/core"
	"go-api/global"
	"go-api/initialize"
	"go-api/tool"

	cmd "github.com/18211167516/go-cmd"
)

// versionCmd represents the version command
var makeCmd = &cmd.Command{
	Use:   "make:model",
	Short: "读取表结构生成Model文件",
	Long:  `读取mysql的表结构转成Model文件`,
	Run: func(Command *cmd.Command, args []string) {
		table, _ := Command.Flags().GetString("table")
		prefix, _ := Command.Flags().GetString("prefix")
		file, _ := Command.Flags().GetString("file")
		global.VP = core.Viper("../static/config/app.toml") //初始化配置
		dsn := initialize.GetMasterDsn()

		vip := core.Viper("./config/cmd.toml")
		config := &tool.T2tConfig{
			StructNameRtrims: vip.GetBool("StructNameRtrims"),
			UcFirstOnly:      vip.GetBool("UcFirstOnly"),
			SavePath:         vip.GetString("SavePath"),
		}
		grom := tool.NewTable2Struct(config)
		grom.
			Table(table).
			Prefix(prefix).
			SavePath(file).
			Dsn(dsn).
			Run()
	},
}

func init() {
	cmd.RootCmd.AddCommand(makeCmd)
	makeCmd.Flags().StringP("table", "t", "", "指定的表名，如果未指定则全部导出")
	makeCmd.Flags().StringP("prefix", "p", "", "表前缀")
	makeCmd.Flags().StringP("file", "f", "", "生成的目录")
}
