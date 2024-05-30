package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/maximus969/users-app/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/", h.createUser)
			users.GET("/", h.getAllUsers)
			users.GET("/:id",h.getUser)
			users.PATCH("/:id", h.updateUser)
			users.DELETE("/:id", h.deleteUser)
		}
	}

	return router
}