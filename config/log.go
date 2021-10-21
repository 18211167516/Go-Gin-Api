package config

type Log struct {
	Level        string //等级
	Formatter    string //格式化模式
	ShowLine     bool   //是否展示行数
	LogInConsole bool   //是否写入文件的同时写入标准输出
	OutFile      bool   //是否写入文件
	LogDir       string //日志文件夹
}
