package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/n0mercy071/post-service-go/pkg/entity"
)

// @Summary createPost
// @Tags posts
// @ID create-post
// @Accept json
// @Produce json
// @Param input body entity.Post true "post info"
// @Success 200 {object} entity.Post
// @Failure 400, 401 errorResponse
// @Security jwtToken
// @Router /api/posts [post]
func (h *Handler) createPost(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		newErrorReponse(c, http.StatusUnauthorized, "userID not found")
		return
	}

	var input entity.Post
	if err := c.BindJSON(&input); err != nil {
		newErrorReponse(c, http.StatusBadRequest, err.Error())
		return
	}

	post, err := h.services.CreatePost(userID.(int), input)
	if err != nil {
		newErrorReponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *Handler) getPostList(c *gin.Context) {
}

func (h *Handler) getPostByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID, _ := c.Get("userID")
	post := entity.Post{ID: id, UserID: userID.(int)}
	c.JSON(http.StatusOK, post)
}

func (h *Handler) updatePost(c *gin.Context) {
}

func (h *Handler) deletePost(c *gin.Context) {
}
