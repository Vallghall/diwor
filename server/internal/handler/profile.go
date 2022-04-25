package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/Valghall/diwor/server/internal/results"
	"gitlab.com/Valghall/diwor/server/internal/users"
	"net/http"
	"strconv"
)

func (h *Handler) userProfile(c *gin.Context) {
	id := c.GetInt(userCtx)
	name := c.GetString(userName)
	username := c.GetString(userLogin)

	recentExps := h.service.Experiment.GetRecentExperiments(id)

	c.JSON(
		http.StatusOK,
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
	name := c.GetString(userName)
	userId := c.GetInt(userCtx)

	expHistory := h.service.Experiment.GetAllUserExperiments(userId)

	c.JSON(
		http.StatusOK,
		struct {
			Name    string
			History []results.ExperimentDigest
		}{
			name,
			expHistory,
		})
}

func (h *Handler) fetchUserExperimentResult(c *gin.Context) {
	name := c.GetString(userName)
	userId := c.GetInt(userCtx)

	alg := c.Query("alg-type")
	sortedId, err := strconv.Atoi(c.Query("sorted-id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.service.Experiment.GetUserExperimentResultBySortedId(alg, userId, sortedId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	switch result.(type) {
	case results.HashAlgorithmsResults:
		c.JSON(http.StatusOK, struct {
			Name    string
			Results results.HashAlgorithmsResults
		}{
			Name:    name,
			Results: result.(results.HashAlgorithmsResults),
		})
	case results.CipherAlgorithmsResults:
		c.JSON(http.StatusOK, struct {
			Name    string
			Results results.CipherAlgorithmsResults
		}{
			Name:    name,
			Results: result.(results.CipherAlgorithmsResults),
		})
	default:
		newErrorResponse(c, http.StatusInternalServerError, "")
	}
}
