package handler

import "github.com/gin-gonic/gin"

type Handler struct{}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		posts := api.Group("/posts")
		{
			posts.POST("", h.createPost)
			posts.GET("", h.getPostList)
			posts.GET("/:id", h.getPostByID)
			posts.PUT("/:id", h.updatePost)
			posts.DELETE("/:id", h.deletePost)
		}
	}

	return router
}
