package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	userName            = "name"
	userLogin           = "userName"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty authorization header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, fmt.Sprintf("invalid authorization header: %v", header))
		return
	}

	userId, err := h.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := h.service.Authorization.GetUserById(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Set(userCtx, userId)
	c.Set(userName, user.Name)
	c.Set(userLogin, user.Username)
}

func (h *Handler) userIdentify(c *gin.Context) {
	tokenCookie, err := c.Request.Cookie("diwor-access-token")
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/auth/login")
		return
	}

	userId, err := h.service.Authorization.ParseToken(tokenCookie.Value)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/auth/login")
		return
	}

	user, err := h.service.Authorization.GetUserById(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Set(userCtx, user.Id)
	c.Set(userName, user.Name)
	c.Set(userLogin, user.Username)
}
