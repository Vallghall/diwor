package handler

import (
	"gitlab.com/Valghall/diwor/server/internal/service"
	"net/http"

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

	router.LoadHTMLFiles("client/build/index.html")
	router.StaticFile("/favicon.ico", "assets/logo/favicon.ico")
	router.Static("/static", "client/build/static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/c/*any", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)

		auth.GET("/logout", h.logout)
	}

	api := router.Group("/api", h.userIdentity)
	{
		prof := api.Group("/profile")
		{
			prof.GET("/", h.userProfile)
			prof.GET("/results/", h.userExperimentHistory)

			prof.GET("/fetch-result", h.fetchUserExperimentResult)
		}

		ex := api.Group("/experiment")
		{
			ex.GET("/", h.indexPage)

			ex.GET("/hashes", h.pickHashingAlgorithms)
			ex.GET("/hash-results", h.hashResults)

			ex.GET("/ciphers", h.pickCipheringAlgorithms)
			ex.GET("/cipher-results", h.cipherResults)

			ex.POST("/start-hash-experiment", h.researchHashAlgorithms)
			ex.POST("/start-cipher-experiment", h.researchCipherAlgorithm)
		}
	}
	return router
}
