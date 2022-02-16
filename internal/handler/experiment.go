package handler

import (
	"github.com/sirupsen/logrus"
	myerr "gitlab.com/Valghall/diwor/internal/errors"
	resulties "gitlab.com/Valghall/diwor/internal/results"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/Valghall/diwor/pkg/experimentator"
)

const (
	HashAlgorithm   = "Алгоритм хэширования"
	CipherAlgorithm = "Алгоритм шифрования"
)

type HashAlgorithmsInput struct {
	Algorithms []string `json:"algorithms"`
}

func (h *Handler) indexPage(c *gin.Context) {

	userInfo, ok := c.Get(userName)
	if !ok {
		logrus.Error(myerr.ErrUserCtxNotFound)
		c.HTML(http.StatusOK, "experiment.gohtml", "Master")
	} else {
		c.HTML(http.StatusOK, "experiment.gohtml", userInfo)
	}
}

func (h *Handler) researchHashAlgorithms(c *gin.Context) {
	userId, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, myerr.ErrUserCtxNotFound.Error())
	}

	var initials HashAlgorithmsInput
	err := c.BindJSON(&initials)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var results resulties.HashAlgorithmsResults
	for _, algorithm := range initials.Algorithms {
		res := h.service.Experiment.ResearchHashingAlgorithm(algorithm, &results)
		if res == nil {
			newErrorResponse(c, http.StatusBadRequest, myerr.ErrInvalidHashAlgorithmInput.Error())
			return
		}
		results.Results = append(results.Results, res)
	}

	h.service.Experiment.SaveResults(userId.(int), HashAlgorithm, results)

	c.JSON(http.StatusOK, results)
}

func (h *Handler) pickHashingAlgorithms(c *gin.Context) {
	name, _ := c.Get(userName)
	c.HTML(http.StatusOK, "hashes.gohtml", name)
}

func (h *Handler) pickCipheringAlgorithms(c *gin.Context) {
	name, _ := c.Get(userName)
	c.HTML(http.StatusOK, "ciphers.gohtml", name)
}

// Deprecated: Need changes
// TODO: Fix this handler
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
