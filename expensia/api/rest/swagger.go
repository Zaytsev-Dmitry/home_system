package rest

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"os"
	"time"
)

func Load(r *gin.Engine) {
	fsys := os.DirFS("third_party/swagger_ui")
	templ := template.Must(template.New("").ParseFS(fsys, "templates/*"))

	r.SetHTMLTemplate(templ)
	r.StaticFS("/public", http.Dir("third_party/swagger_ui"))
	r.StaticFile("/rest/expensia-server-api.yml", "api/rest/expensia-server-api.yml")

	r.GET("/docs", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"SpecVersion": time.Now().Unix(), // кеш-бастер
		})
	})
}
