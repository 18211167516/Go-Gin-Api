package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"go-api/global"

	"github.com/gin-gonic/gin"
)

//默认值方法，使用方法 {{.title|default "123"}} 如果.title有值就显示，反之显示123
func DefaultValueFunc(s interface{}, x interface{}) string {
	if x == nil {
		return fmt.Sprint(s)
	}
	return fmt.Sprint(x)
}

//模板自定义函数
var Funcs = template.FuncMap{
	"default": DefaultValueFunc,
}

func fileRoute(r *gin.Engine) {
	f := global.FS
	templ := template.Must(template.New("").Funcs(Funcs).ParseFS(f, "templates/*/*"))
	r.SetHTMLTemplate(templ)
	// example: /public/static/js/a.js
	r.StaticFS("/public", http.FS(f))
}
