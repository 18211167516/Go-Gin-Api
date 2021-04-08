package routes

import (
	"html/template"
	"net/http"

	"go-api/global"

	"github.com/gin-gonic/gin"
)

func fileRoute(r *gin.Engine) {
	f := global.FS
	templ := template.Must(template.New("").ParseFS(f, "templates/*"))
	r.SetHTMLTemplate(templ)

	// example: /public/static/js/a.js
	r.StaticFS("/public", http.FS(f))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Embed Demo",
		})
	})

	r.GET("/foo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Foo Bar",
		})
	})
}
