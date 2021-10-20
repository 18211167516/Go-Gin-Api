package tests

import (
	"fmt"
	"os"
	"testing"

	//"go-api/core"
	//"go-api/global"

	"github.com/gin-gonic/gin"
)

func setup() {
	//global.VP = core.Viper("../static/config/app.toml")
	//初始化日志
	//初始化DB
	//global.DB = initialize.Gorm()
	gin.SetMode(gin.TestMode)
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("After all tests")
}
func TestMain(m *testing.M) {
	setup()
	fmt.Println("Test begins....")
	code := m.Run() // 如果不加这句，只会执行Main
	teardown()
	os.Exit(code)
}
