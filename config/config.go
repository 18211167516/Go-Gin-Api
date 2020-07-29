package config


import (
	"gopkg.in/ini.v1"
	"log"
	"fmt"
	"os"
	"time"
)

//app struct
type App struct {
	Template string
	PageSize int
	JwtSecret string
	JwtExpiresAt time.Duration
	SigningMethod string
}

var AppSetting = &App{}

//server struct
type Server struct {
	HttpAddress string
	HttpPort int
	ReadTimeout  time.Duration
	WriteTimeout  time.Duration
}

var ServerSetting = &Server{}

//Mysql struct
type Mysql struct {
	MysqlUser string
	MysqlPassword string
	MysqlHost string
	MysqlName string
	MysqlPrefix string
}

var MysqlSetting = &Mysql{}

var (
	Cfg *ini.File
	RunMode string
)

const (
	// AppMode indicates config mode is AppMode.
	AppMode = "app"
	// TestMode indicates config mode is test.
	TestMode = "test"
)

var modeName = AppMode

//set config Mode
func SetConfigMode(value string) {
	switch value {
	case AppMode,"":
		modeName = AppMode
	case TestMode:
		modeName = value
	default:
		panic("config mode unknown: " + value)
	}
}

//get config path Single
func getConfigPath() (path string){
	var pwd,_ = os.Getwd()
	switch modeName {
		case AppMode:
			path = fmt.Sprintf("%s/%s",pwd,"config/app.ini");
		case TestMode:
			path = fmt.Sprintf("%s/../%s",pwd,"config/app.ini");
		default:
			panic("config path unknown")
	}
	return
}

func InitConfig() {
	iniPath := getConfigPath()
	var err error
	Cfg, err = ini.Load(iniPath)
    if err != nil {
        fmt.Printf("Fail to read file: %v", err)
        os.Exit(1)
	}
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
	LoadBasice()
	LoadApp()
	LoadServer()
	LoadDatabase()
}

//加载基础配置
func LoadBasice() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

//加载app配置
func LoadApp() {
	sec, err := Cfg.GetSection("app")
    if err != nil {
        log.Fatalf("Fail to get section 'app': %v", err)
	}
	
	err = sec.MapTo(AppSetting)
	if err != nil {
        log.Fatalf("Cfg.MapTo AppSetting err: %v", err)
    }

	AppSetting.JwtExpiresAt = time.Duration(sec.Key("JWT_EXPIRE_TIME").MustInt(10))*time.Minute
}

//加载http服务配置
func LoadServer() {
	sec, err := Cfg.GetSection("server")
    if err != nil {
        log.Fatalf("Fail to get section 'server': %v", err)
	}
	
	err = sec.MapTo(ServerSetting)
	if err != nil {
        log.Fatalf("Cfg.MapTo ServerSetting err: %v", err)
    }

    ServerSetting.ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
    ServerSetting.WriteTimeout =  time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second 
}

//加载数据库配置
func LoadDatabase() {
	sec, err := Cfg.GetSection("database-mysql")
    if err != nil {
        log.Fatalf("Fail to get section 'app': %v", err)
    }

	err = sec.MapTo(MysqlSetting)
	if err != nil {
        log.Fatalf("Cfg.MapTo MysqlSetting err: %v", err)
	}
}


