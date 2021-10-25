package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.LoadHTMLGlob("web/template/*")
	router.StaticFile("/favicon.ico", "assets/logo/favicon.ico")
	router.Static("/css", "web/static/css")
	router.Static("/image", "assets/image")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.gohtml", nil)
	})

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sing-in", h.signIp)
	}

	ex := router.Group("/experiment")
	{
		ex.GET("/", h.indexPage)
		ex.GET("/results", h.results)
	}
	return router
}
