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

type commandStruct struct {
	Use       string
	CmdName   string
	CmdParent string
}

type serviceStruct struct {
	Package     string //tpl里{{.Package}}
	ModelImport string
	ModelName   string
	ModelStruct string //类似models.SysMenu
}

type controllerStruct struct {
	Package        string //tpl里{{.Package}}
	ControllerName string //控制器名 大写
	ViewName       string //小写
	Name           string //中文名 注释用
	ModelImport    string //类似 app/models
	ModelStruct    string //类似 models.SysMenu
	ServiceImport  string //类似 app/services"
	ServicePackage string //调用的服务层的package
	ServiceName    string //调用的服务名
}

type viewStruct struct {
	Default string
	Desc    string
}

var defaultAppPath = "../app"
var defaultTmpPath = "../templates"

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
	cmdFile, err := tool.CreateFile(cmdFilePath)
	if err != nil {
		return err
	}
	defer cmdFile.Close()

	c := commandStruct{
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

	// 校验并且生成目录
	if err := DirExists(path); err != nil {
		return path, err
	}

	File, err := tool.CreateFile(path + fileName)
	if err != nil {
		return path + fileName, err
	}
	defer File.Close()
	//简析model
	s := serviceStruct{
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

//创建controller
func ControllerCreate(path, name, model, service string) (string, error) {
	//去掉右侧/
	path = strings.TrimRight(path, "/")
	//生成controller
	controller := defaultAppPath + "/controller/" + path
	//按/分割
	consqite := strings.Split(controller, "/")
	//package取倒数第二位
	Package := consqite[len(consqite)-2]
	//视图名称取最后一位
	ViewName := consqite[len(consqite)-1]
	//控制器名称
	ControllerName := tool.CamelCase(ViewName, "_")
	//默认从path匹配 例如 path传 test,那么model默认匹配到 app/models 下的Test
	ModelName := tool.CamelCase(ViewName, "_")
	ModelImport := "app/models"
	ModelStruct := "models." + ModelName

	serviceImport := ""
	servicePackage := "services"
	serviceName := ControllerName

	if patharr := strings.Split(path, "/"); len(patharr) > 1 {
		ModelName := tool.FirstUpper(ViewName)
		pak := patharr[len(patharr)-2]
		ModelImport = "app/models/" + pak
		ModelStruct = pak + "." + ModelName

		serviceImport = "app/services/" + pak
		servicePackage = pak
	}
	//有model的情况下，处理
	if model != "" {
		modelqite := strings.Split(model, "/")
		if len(modelqite) > 1 {
			//取最后一位 并转成首字母大写
			ModelName = tool.FirstUpper(modelqite[len(modelqite)-1])
			ModelImport += "/" + strings.Join(modelqite[:len(modelqite)-1], "/")
			ModelStruct = fmt.Sprintf("%s.%s", modelqite[len(modelqite)-2], ModelName)
		} else {
			//重新设置model name
			ModelName = tool.FirstUpper(model)
			ModelImport = "app/models"
			ModelStruct = "models." + ModelName
		}
	}

	//有service的情况下处理
	if service != "" {
		serviceqite := strings.Split(service, "/")
		if len(serviceqite) > 1 {
			//取最后一位 并转成首字母大写
			serviceName = tool.FirstUpper(serviceqite[len(serviceqite)-1])
			servicePackage = serviceqite[len(serviceqite)-2]
			serviceImport = "app/services/" + strings.Join(serviceqite[:len(serviceqite)-1], "/")
		} else {
			//重新设置model name
			serviceImport = ""
			servicePackage = "services"
			serviceName = tool.FirstUpper(service)
		}
	}

	fileName := fmt.Sprintf("%s.go", controller)

	path = strings.Join(consqite[:len(consqite)-2], "/")
	// 校验并且生成目录
	if err := DirExists(path); err != nil {
		return path, err
	}

	File, err := tool.CreateFile(fileName)
	if err != nil {
		return fileName, err
	}
	defer File.Close()

	s := controllerStruct{
		Package:        Package,
		ControllerName: ControllerName,
		ViewName:       ViewName,
		Name:           name,
		ModelImport:    ModelImport,
		ModelStruct:    ModelStruct,
		ServiceImport:  serviceImport,
		ServicePackage: servicePackage,
		ServiceName:    serviceName,
	}

	controllerTemplate := template.Must(template.New("ser").Parse(tpl.ControllerTemplate()))
	err = controllerTemplate.Execute(File, s)
	if err != nil {
		return fileName, err
	}
	return fileName, nil
}

//创建列表视图
func ViewCreate(path, desc string) error {

	fileArr := strings.Split(path, "/")

	if len(fileArr) != 2 {
		return errors.New("只支持二级目录")
	}

	newpath := defaultTmpPath + "/" + fileArr[0]
	// 校验并且生成目录
	if err := DirExists(newpath); err != nil {
		return err
	}

	fileName := newpath + "/" + fileArr[1] + ".html"
	File, err := tool.CreateFile(fileName)
	if err != nil {
		return err
	}
	defer File.Close()

	v := viewStruct{
		Default: path,
		Desc:    desc,
	}

	controllerTemplate := template.Must(template.New("view").Delims("<!--{", "}-->").Parse(tpl.ViewTemplate()))
	err = controllerTemplate.Execute(File, v)
	if err != nil {
		return err
	}
	return nil
}
