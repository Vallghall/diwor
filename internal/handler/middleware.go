package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(ctx, http.StatusUnauthorized, "empty authorization header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(ctx, http.StatusUnauthorized, "invalid authorization header")
		return
	}

	userId, err := h.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.Set(userCtx, userId)
}
