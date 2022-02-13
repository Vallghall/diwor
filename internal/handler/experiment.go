package handler

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/Valghall/diwor/internal/users"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/Valghall/diwor/pkg/experimentator"
)

func (h *Handler) indexPage(c *gin.Context) {

	userInfo, ok := c.Get(userCtx)
	if !ok {
		logrus.Error("User context not found")
		c.HTML(http.StatusOK, "experiment.gohtml", "Master")
	} else {
		c.HTML(http.StatusOK, "experiment.gohtml", userInfo.(users.User).Name)
	}
}

func (h *Handler) startExperiment(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) results(c *gin.Context) {
	sampleA, sampleB, modeA, modeB :=
		c.Query("sample-a"),
		c.Query("sample-b"),
		c.Query("mode-1"),
		c.Query("mode-2")

	if sampleA == sampleB {
		c.Redirect(http.StatusTemporaryRedirect, "/api/experiment/?reason=equal")
	}
	encryptionResA, _ := experimentator.HoldExperiment(sampleA, sampleB, modeA, modeB)
	c.HTML(http.StatusOK, "results.gohtml", encryptionResA)
}
