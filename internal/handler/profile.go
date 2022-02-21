package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/Valghall/diwor/internal/results"
	"gitlab.com/Valghall/diwor/internal/users"
	"net/http"
)

func (h *Handler) userProfile(c *gin.Context) {
	id := c.GetInt(userCtx)
	name := c.GetString(userName)
	username := c.GetString(userLogin)

	recentExps := h.service.Experiment.GetRecentExperiments(id)

	c.HTML(
		http.StatusOK,
		"profile.gohtml",
		struct {
			users.User
			Digests []results.ExperimentDigest
		}{
			User: users.User{
				Name:     name,
				Username: username,
			},
			Digests: recentExps,
		})
}

func (h *Handler) userExperimentHistory(c *gin.Context) {

}
