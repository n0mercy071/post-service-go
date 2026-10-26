package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n0mercy071/post-service-go/pkg/entity"
)

func (h *Handler) signIn(c *gin.Context) {
}

func (h *Handler) signUp(c *gin.Context) {
	var input entity.User

	if err := c.BindJSON(&input); err != nil {
		newErrorReponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorReponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
