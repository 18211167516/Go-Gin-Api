package main

import (
	"go-api/core"
	"go-api/global"
	"go-api/initialize"
	"go-api/tool"

	cmd "github.com/18211167516/go-cmd"
)

// versionCmd represents the version command
var structCmd = &cmd.Command{
	Use:   "struct",
	Short: "mysql转struct",
	Long:  `读取mysql的表结构转成Model文件`,
	Run: func(Command *cmd.Command, args []string) {

		table, _ := Command.Flags().GetString("table")
		prefix, _ := Command.Flags().GetString("prefix")
		file, _ := Command.Flags().GetString("file")

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
		// 指定某个表,如果不指定,则默认全部表都迁移
		/* Table(*table).
		// 表前缀
		Prefix(*prefix).
		// 是否添加json tag
		EnableJsonTag(*enableJsonTag).
		// 生成struct的包名(默认为空的话, 则取名为: package model)
		PackageName(*packageName).
		// tag字段的key值,默认是gorm
		TagKey(*tagKey).
		// 是否添加结构体方法获取表名
		RealNameMethod(*realNameMethod).
		// 生成的结构体保存路径
		SavePath(*file).
		// 数据库dsn
		Dsn(*dsn). */
		// 执行
		//Run()
	},
}

func init() {
	cmd.RootCmd.AddCommand(structCmd)
	global.VP = core.Viper("../static/config/app.toml") //初始化配置
	structCmd.Flags().StringP("table", "t", "", "指定的表名，如果未指定则全部导出")
	structCmd.Flags().StringP("prefix", "p", "", "表前缀")
	structCmd.Flags().StringP("file", "f", "", "生成的目录")
}
