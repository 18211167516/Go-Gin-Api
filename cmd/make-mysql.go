package main

import (
	"go-api/cmd/initdata"
	"go-api/core"
	"go-api/global"
	"go-api/initialize"

	cmd "github.com/18211167516/go-cmd"
)

// versionCmd represents the version command
var makeMysqlCmd = &cmd.Command{
	Use:   "make:mysql",
	Short: "读取表结构生成Model文件",
	Long:  `读取mysql的表结构转成Model文件`,
	Run: func(Command *cmd.Command, args []string) {
		table, _ := Command.Flags().GetString("table")
		global.VP = core.Viper("../static/config/app.toml") //初始化配置
		global.VP.Set("mysql.global.LogMode", "Warn")
		global.DB = initialize.Gorm() //初始化DB
		initdata.AutoMigrate(global.DB, table)
	},
}

func init() {
	cmd.RootCmd.AddCommand(makeMysqlCmd)
	makeMysqlCmd.Flags().StringP("table", "t", "", "指定的表名,必须指定")
	makeMysqlCmd.MarkFlagRequired("table")
}
