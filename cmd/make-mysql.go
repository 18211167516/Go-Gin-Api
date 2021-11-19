package main

import (
	"go-api/cmd/util"
	"go-api/core"
	"go-api/global"
	"go-api/initialize"

	cmd "github.com/18211167516/go-cmd"
)

// versionCmd represents the version command
var makeMysqlCmd = &cmd.Command{
	Use:   "make:mysql",
	Short: "读取model生成表",
	Long:  `读取model生成数据库表结构`,
	Example: `
读取 Test要看init方法的key
默认生成到app/models下
./cmd.exe make:model -t=Test
生成到app/models/test下
./cmd.exe make:model -t=Test -f=test
`,
	Run: func(Command *cmd.Command, args []string) {
		table, _ := Command.Flags().GetString("table")
		global.VP = core.Viper("../static/config/app.toml") //初始化配置
		global.VP.Set("mysql.global.LogMode", "Warn")
		global.DB = initialize.Gorm() //初始化DB
		util.AutoMigrate(global.DB, table)
	},
}

func init() {
	cmd.RootCmd.AddCommand(makeMysqlCmd)
	makeMysqlCmd.Flags().StringP("table", "t", "", "指定的表名,必须指定")
	makeMysqlCmd.MarkFlagRequired("table")
}
