package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n0mercy071/post-service-go/pkg/entity"
)

type signInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorReponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorReponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"access_token": token,
	})
}

func (h *Handler) signUp(c *gin.Context) {
	var input entity.User

	if err := c.BindJSON(&input); err != nil {
		newErrorReponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateUser(input)
	if err != nil {
		newErrorReponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"id": id,
	})
}
