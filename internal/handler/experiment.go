package handler

import (
	"github.com/sirupsen/logrus"
	myerr "gitlab.com/Valghall/diwor/internal/errors"
	"gitlab.com/Valghall/diwor/internal/results"
	"net/http"

	"github.com/gin-gonic/gin"
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
		newErrorResponse(c, http.StatusUnauthorized, myerr.ErrUserCtxNotFound.Error())
		return
	}

	var initials HashAlgorithmsInput
	err := c.BindJSON(&initials)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var algResults results.HashAlgorithmsResults
	for _, algorithm := range initials.Algorithms {
		res := h.service.Experiment.ResearchHashingAlgorithm(algorithm, &algResults)
		algResults.Results = append(algResults.Results, res)
	}

	h.service.Experiment.SaveResults(userId.(int), HashAlgorithm, algResults)

	c.JSON(http.StatusOK, algResults)
}

func (h *Handler) pickHashingAlgorithms(c *gin.Context) {
	name, _ := c.Get(userName)
	c.HTML(http.StatusOK, "hashes.gohtml", name)
}

func (h *Handler) pickCipheringAlgorithms(c *gin.Context) {
	name, _ := c.Get(userName)
	c.HTML(http.StatusOK, "ciphers.gohtml", name)
}

func (h *Handler) hashResults(c *gin.Context) {
	name := c.GetString(userName)
	userId := c.GetInt(userCtx)

	res := h.service.Experiment.GetLastExperimentResults(userId)

	c.HTML(http.StatusOK, "hash-results.gohtml", struct {
		Name    string
		Results results.HashAlgorithmsResults
	}{
		Name:    name,
		Results: res,
	})
}
