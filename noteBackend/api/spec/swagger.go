package generatedApi

import (
	"embed"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

//go:embed assets templates openapi
var embeddedFiles embed.FS

func Load(r *gin.Engine) {
	templ := template.Must(template.New("").ParseFS(embeddedFiles, "templates/*"))
	r.SetHTMLTemplate(templ)
	r.StaticFS("/public", http.FS(embeddedFiles))

	r.GET("/docs", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})
}
