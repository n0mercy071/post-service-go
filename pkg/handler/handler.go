package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/n0mercy071/post-service-go/pkg/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/n0mercy071/post-service-go/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("sign-in", h.signIn)
		auth.POST("sign-up", h.signUp)
	}

	api := router.Group("/api", h.userIdentity)
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
