package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/n0mercy071/post-service-go/pkg/domain/entity"
)

func (h *Handler) createPost(c *gin.Context) {
}

func (h *Handler) getPostList(c *gin.Context) {
}

func (h *Handler) getPostByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	post := entity.Post{ID: id}
	c.JSON(http.StatusOK, post)
}

func (h *Handler) updatePost(c *gin.Context) {
}

func (h *Handler) deletePost(c *gin.Context) {
}
