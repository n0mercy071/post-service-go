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

// @Summary signIn
// @Tags auth
// @ID create-user
// @Accept json
// @Produce json
// @Param input body signInInput true "sign in input"
// @Success 200 {integer} integer 1
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
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

// @Summary signUp
// @Tags auth
// @ID sign-up
// @Accept json
// @Produce json
// @Param input body entity.User true "user info"
// @Success 200 {integer} integer 1
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
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
