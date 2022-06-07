package handler

import (
	"github.com/sirupsen/logrus"
	myerr "gitlab.com/Valghall/diwor/server/internal/errors"
	"gitlab.com/Valghall/diwor/server/internal/plotconfig"
	"gitlab.com/Valghall/diwor/server/internal/results"
	"net/http"
	"os/exec"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	HashAlgorithm   = "Алгоритм хэширования"
	CipherAlgorithm = "Алгоритм шифрования"
)

type AlgorithmsInput struct {
	Algorithms []string `json:"algorithms"`
	plotconfig.Config
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
	userId := c.GetInt(userCtx)

	var initials AlgorithmsInput
	err := c.BindJSON(&initials)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var algResults results.HashAlgorithmsResults
	loc, _ := time.LoadLocation("Europe/Moscow")
	begin := time.Now().In(loc)

	for _, algorithm := range initials.Algorithms {
		res := h.service.Experiment.ResearchHashingAlgorithm(algorithm, initials.Config)
		algResults.Results = append(algResults.Results, res)
	}
	end := time.Now().In(loc)

	algResults.StartedAt = begin
	algResults.FinishedAt = end

	algResults.OS = runtime.GOOS
	algResults.Arch = runtime.GOARCH
	proc, err := getProcessorInfo()
	if err == nil {
		algResults.Processor = string(proc)
	}

	h.service.Experiment.SaveResults(userId, HashAlgorithm, algResults)

	c.JSON(http.StatusOK, algResults)
}

func (h *Handler) researchCipherAlgorithm(c *gin.Context) {
	userId := c.GetInt(userCtx)

	var initials AlgorithmsInput
	err := c.BindJSON(&initials)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var algResults results.CipherAlgorithmsResults
	loc, _ := time.LoadLocation("Europe/Moscow")
	begin := time.Now().In(loc)
	for _, algorithm := range initials.Algorithms {
		res := h.service.Experiment.ResearchCipheringAlgorithm(algorithm, initials.Config)
		algResults.Results = append(algResults.Results, res)
	}
	end := time.Now().In(loc)

	algResults.StartedAt = begin
	algResults.FinishedAt = end

	algResults.OS = runtime.GOOS
	algResults.Arch = runtime.GOARCH

	h.service.Experiment.SaveResults(userId, CipherAlgorithm, algResults)

	c.JSON(http.StatusOK, algResults)
}

func (h *Handler) pickHashingAlgorithms(c *gin.Context) {
	name := c.GetString(userName)
	c.HTML(http.StatusOK, "hashes.gohtml", name)
}

func (h *Handler) pickCipheringAlgorithms(c *gin.Context) {
	name := c.GetString(userName)
	c.HTML(http.StatusOK, "ciphers.gohtml", name)
}

func (h *Handler) hashResults(c *gin.Context) {
	name := c.GetString(userName)
	userId := c.GetInt(userCtx)

	res := h.service.Experiment.GetLastHashExperimentResults(userId)

	c.HTML(http.StatusOK, "hash-results.gohtml", struct {
		Name    string
		Results results.HashAlgorithmsResults
	}{
		Name:    name,
		Results: res,
	})
}

func (h *Handler) cipherResults(c *gin.Context) {
	name := c.GetString(userName)
	userId := c.GetInt(userCtx)

	res := h.service.Experiment.GetLastCipherExperimentResults(userId)

	c.HTML(http.StatusOK, "cipher-results.gohtml", struct {
		Name    string
		Results results.CipherAlgorithmsResults
	}{
		Name:    name,
		Results: res,
	})
}

func getProcessorInfo() (out []byte, err error) {
	if runtime.GOOS == "windows" {
		out, err = exec.Command("wmic", "cpu", "get", "name").Output()
		return
	}

	grep := exec.Command("grep", "'Model name'")
	lscpu := exec.Command("lscpu")

	pipe, _ := lscpu.StdoutPipe()
	defer pipe.Close()

	grep.Stdin = pipe
	lscpu.Start()

	out, err = grep.Output()
	return
}
