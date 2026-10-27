package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/n0mercy071/post-service-go/pkg/entity"
)

func (h *Handler) createPost(c *gin.Context) {
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
