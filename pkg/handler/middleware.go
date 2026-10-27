package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		newErrorReponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorReponse(c, http.StatusUnauthorized, "incorrect auth header")
		return
	}

	userID, err := h.services.ParseToken(headerParts[1])
	if err != nil {
		newErrorReponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("userID", userID)
}
