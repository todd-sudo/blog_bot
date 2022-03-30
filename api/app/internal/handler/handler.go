package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/todd-sudo/blog_bot/api/internal/service"
	"github.com/todd-sudo/blog_bot/api/pkg/logging"
)

type Handler struct {
	service *service.Service
	log     logging.Logger
}

func NewHandler(service *service.Service, log logging.Logger) *Handler {
	return &Handler{
		service: service,
		log:     log,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()

	blogRoutes := r.Group("api/posts")
	{
		blogRoutes.GET("/", h.AllPost)
		blogRoutes.POST("/", h.InsertPost)
		blogRoutes.DELETE("/:id", h.DeletePost)
		blogRoutes.GET("/:id", h.DeletePost)
	}

	categoryRoutes := r.Group("api/categories")
	{
		categoryRoutes.GET("/", h.AllCategory)
		categoryRoutes.POST("/", h.InsertCategory)
		categoryRoutes.DELETE("/:id", h.DeleteCategory)
	}

	userRoutes := r.Group("api/user")
	{
		userRoutes.GET("/profile", h.ProfileUser)
		userRoutes.POST("/create-user", h.CreateUser)
	}

	return r
}
