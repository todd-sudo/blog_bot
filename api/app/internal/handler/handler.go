package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/todd-sudo/blog_bot/api/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()

	blogRoutes := r.Group("api/posts", nil)
	{
		blogRoutes.GET("/", h.AllPost)
		blogRoutes.POST("/", h.InsertPost)
		blogRoutes.GET("/:id", h.FindByIDPost)
		blogRoutes.PUT("/:id", h.UpdatePost)
		blogRoutes.DELETE("/:id", h.DeletePost)
	}

	return r
}
