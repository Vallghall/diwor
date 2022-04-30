package handler

import (
	"database/sql"
	myerr "gitlab.com/Valghall/diwor/server/internal/errors"
	"gitlab.com/Valghall/diwor/server/internal/users"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const RefreshTokenCookieName = "not-a-token"

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

	accessToken, refreshToken, err := h.service.Authorization.GenerateTokenPair(input.Username, input.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			newErrorResponse(c, http.StatusBadRequest, myerr.ErrInvalidLoginOrPassword.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     RefreshTokenCookieName,
		Value:    refreshToken,
		Expires:  time.Now().Add(12 * time.Hour),
		Path:     "/",
		HttpOnly: true,
	})

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": accessToken,
	})
}

func (h *Handler) refresh(c *gin.Context) {
	refreshToken, err := c.Request.Cookie(RefreshTokenCookieName)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "no refresh token")
		return
	}

	userId, err := h.service.ParseToken(refreshToken.Value)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newAccessToken, newRefreshToken, err := h.service.Authorization.RegenerateTokenPair(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			newErrorResponse(c, http.StatusBadRequest, myerr.ErrInvalidLoginOrPassword.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     RefreshTokenCookieName,
		Value:    newRefreshToken,
		Expires:  time.Now().Add(12 * time.Hour),
		Path:     "/",
		HttpOnly: true,
	})

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": newAccessToken,
	})
}

func (h *Handler) logout(c *gin.Context) {
	cookie := &http.Cookie{
		Name:    RefreshTokenCookieName,
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),

		HttpOnly: true,
	}

	http.SetCookie(c.Writer, cookie)

	c.JSON(http.StatusOK, map[string]string{
		"message": "logged out",
	})
}
