package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/Valghall/diwor/pkg/experimentator"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("web/template/*")
	r.StaticFile("/favicon.ico", "./web/static/favicon.ico")
	r.Static("/css", "web/static/css")
	r.Static("/image", "./web/static/image")

	r.GET("/", func(c *gin.Context) {
		var warning string

		if _, ok := c.GetQuery("reason"); ok {
			warning = "Алгоритмы не должны совпадать!"
		}

		c.HTML(http.StatusOK, "index.gohtml", warning)
	})

	r.GET("/experiment", func(c *gin.Context) {
		sampleA, sampleB, modeA, modeB :=
			c.Query("sample-a"),
			c.Query("sample-b"),
			c.Query("mode-1"),
			c.Query("mode-2")

		if sampleA == sampleB {
			c.Redirect(http.StatusTemporaryRedirect, "/?reason=equal")
		}
		encryptionResA, _ := experimentator.HoldExperiment(sampleA, sampleB, modeA, modeB)
		c.HTML(http.StatusOK, "experiment.gohtml", encryptionResA)
	})

	log.Fatalln(r.Run())
}
