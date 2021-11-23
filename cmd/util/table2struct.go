package util

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var (
	tmpModel1 = `type %s struct {
	Model
`
	tmpModel2 = `type %s struct {
`
	tmpModel3 = `type %s struct {
	models.Model
`
	funcTmp1 = `
/*
func init() {
	AutoMigratFunc["%s"] = func() interface{} {
		return %s{}
	}
}
*/
`
	funcTmp2 = `
/*
func init() {
	models.AutoMigratFunc["%s"] = func() interface{} {
		return %s{}
	}
}
*/
`
)

//map for converting mysql type to golang types
var typeForMysqlToGo = map[string]string{
	"int":                "int",
	"integer":            "int",
	"tinyint":            "int",
	"smallint":           "int",
	"mediumint":          "int",
	"bigint":             "int64",
	"int unsigned":       "int64",
	"integer unsigned":   "int64",
	"tinyint unsigned":   "int",
	"smallint unsigned":  "int",
	"mediumint unsigned": "int64",
	"bigint unsigned":    "int64",
	"bit":                "int64",
	"bool":               "bool",
	"enum":               "string",
	"set":                "string",
	"varchar":            "string",
	"char":               "string",
	"tinytext":           "string",
	"mediumtext":         "string",
	"text":               "string",
	"longtext":           "string",
	"blob":               "string",
	"tinyblob":           "string",
	"mediumblob":         "string",
	"longblob":           "string",
	"date":               "XTime", // time.Time or string
	"datetime":           "XTime", // time.Time or string
	"timestamp":          "XTime", // time.Time or string
	"time":               "XTime", // time.Time or string
	"float":              "float64",
	"double":             "float64",
	"decimal":            "float64",
	"binary":             "string",
	"varbinary":          "string",
}

type Table2Struct struct {
	dsn         string
	db          *sql.DB
	table       string
	prefix      string
	config      *T2tConfig
	packageName string // 生成struct的包名(默认为空的话, 则取名为: package model)
	tagKey      string // tag字段的key值,默认是orm
}

type T2tConfig struct {
	SavePath         string
	UcFirstOnly      bool // 字段首字母大写的同时, 是否要把其他字母转换为小写,默认false不转换
	StructNameRtrims bool //默认去除表名尾部s字符串
}

type column struct {
	ColumnName    string
	Type          string
	Nullable      string
	TableName     string
	ColumnComment string
	Tag           string
}

type count struct {
	TableName string
	Cnt       int
	Tmp       string
}

var packageName = `package %s 

`

func NewTable2Struct(c *T2tConfig) *Table2Struct {
	return &Table2Struct{
		tagKey: "gorm",
		config: c,
	}
}

func (t *Table2Struct) Dsn(d string) *Table2Struct {
	t.dsn = d
	return t
}

func (t *Table2Struct) TagKey(r string) *Table2Struct {
	if r != "" {
		t.tagKey = r
	}
	return t
}

func (t *Table2Struct) SavePath(p string) *Table2Struct {
	if p != "" {
		t.config.SavePath = t.config.SavePath + p
	}
	return t
}

func (t *Table2Struct) DB(d *sql.DB) *Table2Struct {
	t.db = d
	return t
}

func (t *Table2Struct) Table(tab string) *Table2Struct {
	t.table = tab
	return t
}

func (t *Table2Struct) Prefix(p string) *Table2Struct {
	t.prefix = p
	return t
}

func (t *Table2Struct) Config(c *T2tConfig) *Table2Struct {
	t.config = c
	return t
}

//单独执行生成Model
func (t *Table2Struct) single_run(tmp string, tableRealName string, item []column) {
	tableName := tableRealName
	// 去除前缀
	if t.prefix != "" {
		tableName = tableRealName[len(t.prefix):]
	}
	structName := t.Rtrim(tableName)
	depth := 1
	structContent := fmt.Sprintf(tmp, structName)
	for _, v := range item {
		// 字段注释
		var clumnComment string
		if v.ColumnComment != "" {
			clumnComment = fmt.Sprintf(" // %s", v.ColumnComment)
		}
		structContent += fmt.Sprintf("%s%s %s %s%s\n",
			tab(depth), v.ColumnName, v.Type, v.Tag, clumnComment)
	}
	structContent += tab(depth-1) + "}\n\n"

	// 添加 method 获取真实表名
	//structContent += fmt.Sprintf(funcTmp1, structName, tableRealName)
	//if t.realNameMethod != "" {
	/* structContent += fmt.Sprintf("func (%s) %s() string {\n",
		structName, t.realNameMethod)
	structContent += fmt.Sprintf("%sreturn \"%s\"\n",
		tab(depth), tableRealName)
	structContent += "}\n\n" */
	//}

	var importContent string
	if strings.Contains(structContent, "gorm.DeletedAt") {
		importContent = "import \"gorm.io/gorm\"\n\n"
	}

	fucTmp := ""
	if strings.Contains(structContent, "models") {
		importContent = "import \"go-api/app/models\"\n\n"
		fucTmp = funcTmp2
	} else {
		fucTmp = funcTmp1
	}

	structContent += fmt.Sprintf(fucTmp, structName, structName)
	t.saveFile(tableName, importContent, structContent)
}

func createDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func DirExists(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}

	if os.IsNotExist(err) {
		return createDir(path)
	}
	return nil
}

func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

//保存单个Model文件
func (t *Table2Struct) saveFile(tableName, importContent, structContent string) {
	// 是否指定保存路径
	savePath := fmt.Sprintf("%s/%s.go", t.config.SavePath, tableName)

	if isExits := FileExist(savePath); isExits {
		log.Printf("表【%s】文件已存在!!!\r\n", tableName)
		return
	}

	f, err := os.Create(savePath)
	if err != nil {
		log.Printf("表【%s】写入文件失败 【错误】：%s\r\n", tableName, err)
		return
	}
	defer f.Close()

	f.WriteString(t.packageName + importContent + structContent)

	cmd := exec.Command("gofmt", "-w", savePath)
	cmd.Run()

	log.Printf("表【%s】生成成功 \r\n", tableName)
}

//执行生成model动作
//1、校验mysql链接
//2、检测文件夹是否存在
//3、获取表信息
//4、循环执行生成model
func (t *Table2Struct) Run() {

	// 链接mysql, 获取db对象
	t.dialMysql()

	// 校验并且生成目录
	if err := DirExists(t.config.SavePath); err != nil {
		log.Printf("目录%s 错误 %s", t.config.SavePath, err)
		return
	}

	//获取目录最后一个元素
	if base := path.Base(t.config.SavePath); base == ".." || base == "." || base == "/" {
		log.Printf("Model目录非法，不能是%s", t.config.SavePath)
		return
	} else {
		t.packageName = fmt.Sprintf(packageName, base)
	}

	log.Printf("生成目录%s", t.config.SavePath)

	// 获取表和字段的shcema
	tableColumns, tableCount, err := t.getColumns()
	if err != nil {
		log.Println(err)
		return
	}

	// 循环执行
	for tableRealName, item := range tableColumns {
		t.single_run(tableCount[tableRealName].Tmp, tableRealName, item)
	}

	log.Println("gen model finish!!!")
}

func (t *Table2Struct) dialMysql() {
	if t.db == nil {
		if t.dsn == "" {
			log.Println("dsn数据库配置缺失")
			return
		}
		if db, err := sql.Open("mysql", t.dsn); err != nil {
			log.Println(err)
			return
		} else {
			t.db = db
		}

	}
	return
}

func (t *Table2Struct) getCountTable(table ...string) (tableCount map[string]count, err error) {
	tableCount = make(map[string]count)
	var countSql = `SELECT count(*) as Cnt,TABLE_NAME
	FROM information_schema.COLUMNS 
	WHERE table_schema = DATABASE() and column_name IN ('id','created_at','updated_at','deleted_at') `
	if t.table != "" {
		countSql += fmt.Sprintf(" AND TABLE_NAME = '%s'", t.prefix+t.table)
	}
	// sql排序
	countSql += " group by TABLE_NAME order by TABLE_NAME asc, ORDINAL_POSITION asc"
	rows, err := t.db.Query(countSql)
	if err != nil {
		log.Println("Error reading table information count: ", err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		col := count{}
		err = rows.Scan(&col.Cnt, &col.TableName)

		if err != nil {
			log.Println(err.Error())
			return
		}

		if col.Cnt >= 4 {
			col.Tmp = tmpModel1
			if !strings.Contains(t.packageName, "models") {
				col.Tmp = tmpModel3
			}
		} else {
			col.Tmp = tmpModel2
		}

		tableCount[col.TableName] = col
	}
	return
}

// Function for fetching schema definition of passed table
func (t *Table2Struct) getColumns(table ...string) (tableColumns map[string][]column, tableCount map[string]count, err error) {
	// 根据设置,判断是否要把 date 相关字段替换为 string
	tableColumns = make(map[string][]column)
	tableCount, err = t.getCountTable()

	if err != nil {
		log.Println("Error reading table : ", err.Error())
		return
	}

	// sql
	var sqlStr = `SELECT COLUMN_NAME,DATA_TYPE,IS_NULLABLE,TABLE_NAME,COLUMN_COMMENT
		FROM information_schema.COLUMNS 
		WHERE table_schema = DATABASE()`
	// 是否指定了具体的table
	if t.table != "" {
		sqlStr += fmt.Sprintf(" AND TABLE_NAME = '%s'", t.prefix+t.table)
	}
	// sql排序
	sqlStr += " order by TABLE_NAME asc, ORDINAL_POSITION asc"

	rows, err := t.db.Query(sqlStr)
	if err != nil {
		log.Println("Error reading table information: ", err.Error())
		return
	}

	defer rows.Close()
	for rows.Next() {
		col := column{}
		err = rows.Scan(&col.ColumnName, &col.Type, &col.Nullable, &col.TableName, &col.ColumnComment)

		if err != nil {
			log.Println(err.Error())
			return
		}

		if tableCount[col.TableName].Cnt == 4 && (col.ColumnName == "id" || col.ColumnName == "created_at" || col.ColumnName == "updated_at" || col.ColumnName == "deleted_at") {
			continue
		}
		var formTag, tag string

		col.Tag = col.ColumnName
		col.ColumnName = t.camelCase(col.ColumnName)
		col.Type = typeForMysqlToGo[col.Type]
		if col.ColumnName == "DeletedAt" {
			col.Type = "gorm.DeletedAt"
		}
		col.Tag = strings.ToLower(col.Tag)

		//jsonTag = t.camelCase(col.Tag)
		if col.Tag == "id" {
			formTag = fmt.Sprintf(" uri:\"%s\"", "id")
		} else {
			formTag = fmt.Sprintf(" form:\"%s\"", col.Tag)
		}

		tag = fmt.Sprintf("column:%s;", col.Tag)

		if col.ColumnComment != "" {
			tag += fmt.Sprintf("comment:%s", col.ColumnComment)
			col.Tag = fmt.Sprintf("`desc:\"%s\" %s:\"%s\" json:\"%s\" %s`", col.ColumnComment, t.tagKey, tag, col.Tag, formTag)
		} else {
			col.Tag = fmt.Sprintf("`%s:\"%s\" json:\"%s\" %s`", t.tagKey, tag, col.Tag, formTag)
		}

		if col.ColumnName == "Id" {
			col.ColumnName = "ID"
		}

		if _, ok := tableColumns[col.TableName]; !ok {
			tableColumns[col.TableName] = []column{}
		}
		tableColumns[col.TableName] = append(tableColumns[col.TableName], col)
	}

	if len(tableColumns) == 0 {
		log.Println("Error reading table null ")
		return
	}
	return
}

//右裁剪s字符串
func (t *Table2Struct) Rtrim(str string) string {
	text := t.camelCase(str)
	if t.config.StructNameRtrims && text[len(text)-1] == 's' {
		text = text[0 : len(text)-1]
	}
	return text
}

func (t *Table2Struct) camelCase(str string) string {
	// 是否有表前缀, 设置了就先去除表前缀
	if t.prefix != "" {
		str = strings.Replace(str, t.prefix, "", 1)
	}
	var text string
	//for _, p := range strings.Split(name, "_") {
	for _, p := range strings.Split(str, "_") {
		// 字段首字母大写的同时, 是否要把其他字母转换为小写

		switch length := len(p); length {
		case 0:
		case 1:
			text += strings.ToUpper(p[0:1])
		default:
			// 字符长度大于1时
			if t.config.UcFirstOnly == true {
				text += strings.ToUpper(p[0:1]) + strings.ToLower(p[1:])
			} else {
				text += strings.ToUpper(p[0:1]) + p[1:]
			}
		}
	}

	return text
}
func tab(depth int) string {
	return strings.Repeat("\t", depth)
}
