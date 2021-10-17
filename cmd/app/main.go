package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("web/template/*.gohtml")
	r.StaticFile("/favicon.ico", "./web/static/favicon.ico")
	r.StaticFile("/altera", "./web/static/altera.gif")
	r.StaticFile("/css/style.css", "web/template/css/style.css")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.gohtml", nil)
	})
	log.Fatalln(r.Run())
}
