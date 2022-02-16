package handler

import (
	"database/sql"
	"net/http"
	"time"

	myerr "gitlab.com/Valghall/diwor/internal/errors"

	"github.com/gin-gonic/gin"
	"gitlab.com/Valghall/diwor/internal/users"
)

func (h *Handler) signUp(c *gin.Context) {
	var input users.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	ok, err := h.service.Authorization.ValidateUserCredentials(input)
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		if err.Error() == myerr.ErrUsernameAlreadyExists.Error() {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			newErrorResponse(c, http.StatusBadRequest, myerr.ErrInvalidLoginOrPassword.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "diwor-access-token",
		Value:    token,
		Expires:  time.Now().Add(12 * time.Hour),
		Path:     "/",
		HttpOnly: true,
	})

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) logIn(c *gin.Context) {
	c.HTML(http.StatusOK, "login.gohtml", nil)
}

func (h *Handler) register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.gohtml", nil)
}

func (h *Handler) logout(c *gin.Context) {
	cookie := &http.Cookie{
		Name:    "diwor-access-token",
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),

		HttpOnly: true,
	}

	http.SetCookie(c.Writer, cookie)

	c.Redirect(http.StatusTemporaryRedirect, "/auth/login")
}
