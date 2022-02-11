package handler

import (
	"net/http"

	"gitlab.com/Valghall/diwor/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Services
}

func NewHandler(service *service.Services) *Handler {
	return &Handler{service: service}
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
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{

		ex := api.Group("/experiment")
		{
			ex.GET("/", h.indexPage)
			ex.POST("/start", h.startExperiment)
			ex.GET("/results", h.results)
		}
	}
	return router
}
