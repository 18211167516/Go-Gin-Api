package core

import (
	"bytes"
	"flag"
	"fmt"
	"go-api/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
			config = "static/config/app.toml"
			fmt.Printf("您正在使用config的默认值,config的路径为%v\n", config)
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)

	err := ReadInConfig(v, config)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}
	v.WatchConfig()

	//监听文件修改
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.CF); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&global.CF); err != nil {
		fmt.Println("config Unmarshal err:", err)
	}

	return v
}

//读取配置文件 (支持Embed打包包含在二进制文件)
func ReadInConfig(v *viper.Viper, config string) error {
	if err := v.ReadInConfig(); err == nil {
		return err
	}

	file, err := global.FS.ReadFile(config)

	if err != nil {
		return err
	}

	if err := v.ReadConfig(bytes.NewReader(file)); err != nil {
		return err
	}

	fmt.Printf("您正在使用Emabed,config的路径为%v\n", config)
	return nil
}
