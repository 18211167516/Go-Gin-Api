package util

import (
	"errors"
	"fmt"
	"go-api/app/models"
	"go-api/cmd/tpl"
	"go-api/tool"
	"log"
	"os"
	"strings"
	"text/template"

	"gorm.io/gorm"
)

type command struct {
	Use       string
	CmdName   string
	CmdParent string
}

type service struct {
	Package     string //tpl里{{.Package}}
	ModelImport string
	ModelName   string
	ModelStruct string //类似models.SysMenu
}

var defaultAppPath = "../app/"

//通过model生成mysql
func AutoMigrate(db *gorm.DB, table string) {

	if db.Migrator().HasTable(table) {
		log.Printf("[make:mysql]-->数据表【%s】已存在\n", table)
		os.Exit(0)
	}

	if value, ok := models.AutoMigratFunc[table]; !ok {
		log.Printf("make:mysql-->数据表【%s】没有定义model层init方法初始化struct\n", table)
		os.Exit(0)
	} else {
		if err := db.AutoMigrate(value()); err != nil {
			log.Printf("[make:mysql]-->生成数据表【%s】失败,err: %v\n", table, err)
			os.Exit(0)
		}
	}

	log.Println("[make:mysql]-->生成数据表【", table, "】成功")

}

//创建命令文件
func CommandCreate(use, fileName, cmdName, cmdParent string) error {
	cmdFilePath := fmt.Sprintf("%s.go", fileName)
	if success, _ := tool.PathExists(cmdFilePath); success {
		return errors.New("文件已存在")
	}
	cmdFile, err := os.Create(cmdFilePath)
	if err != nil {
		return err
	}
	defer cmdFile.Close()

	c := command{
		Use:       use,
		CmdName:   cmdName,
		CmdParent: cmdParent,
	}
	commandTemplate := template.Must(template.New("sub").Parse(tpl.AddCommandTemplate()))
	err = commandTemplate.Execute(cmdFile, c)
	if err != nil {
		return err
	}
	return nil
}

//创建service文件
func ServiceCreate(model, fileName, path string) (string, error) {
	path = defaultAppPath + "/services/" + path + "/"
	modelsqite := strings.Split(model, "/")

	ModelImport := "app/models"
	ModelStruct := ""
	ModelName := modelsqite[len(modelsqite)-1]
	if len(modelsqite) > 1 {
		ModelImport += "/" + strings.Join(modelsqite[:len(modelsqite)-1], "/")
		ModelStruct = fmt.Sprintf("%s.%s", modelsqite[len(modelsqite)-2], ModelName)
	} else {
		ModelStruct = "models." + ModelName
	}

	fmt.Println(ModelStruct)
	PathArr := strings.Split(strings.TrimRight(path, "/"), "/")
	Package := PathArr[len(PathArr)-1]

	//打开文件
	if fileName == "" {
		fileName = fmt.Sprintf("%sService.go", ModelName)
	} else {
		fileName = fmt.Sprintf("%s.go", fileName)
	}

	if success, _ := tool.PathExists(path + fileName); success {
		return path + fileName, errors.New("文件已存在")
	}

	// 校验并且生成目录
	if err := DirExists(path); err != nil {
		return path, err
	}

	File, err := os.Create(path + fileName)
	if err != nil {
		return path + fileName, err
	}
	defer File.Close()
	//简析model
	s := service{
		Package:     Package, //tpl里{{.Package}}
		ModelImport: ModelImport,
		ModelName:   ModelName,
		ModelStruct: ModelStruct, //类似models.SysMenu
	}

	serviceTemplate := template.Must(template.New("ser").Parse(tpl.ServiceTemplate()))
	err = serviceTemplate.Execute(File, s)
	if err != nil {
		return path + fileName, err
	}
	return path + fileName, nil
}
