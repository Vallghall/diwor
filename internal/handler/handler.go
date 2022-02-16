package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"

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
	router.Static("/js", "web/static/js")
	router.Static("/image", "assets/image")

	router.GET("/", func(c *gin.Context) {
		userInfo, ok := c.Get(userCtx)
		if !ok {
			logrus.Error("User context not found")
		}
		c.HTML(http.StatusOK, "index.gohtml", userInfo)
	})

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)

		auth.GET("/logout", h.logout)
		auth.GET("/login", h.logIn)
		auth.GET("/register", h.register)
	}

	api := router.Group("/api", h.userIdentify)
	{

		ex := api.Group("/experiment")
		{
			ex.GET("/", h.indexPage)
			ex.GET("/hashes", h.pickHashingAlgorithms)
			ex.GET("/ciphers", h.pickCipheringAlgorithms)
			ex.POST("/start-hash-experiment", h.researchHashAlgorithms)
		}
	}
	return router
}
