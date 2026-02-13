package handler

import (
	"github.com/VSBrilyakov/subsApp/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "github.com/VSBrilyakov/subsApp/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")

	{
		v1 := api.Group("/v1")
		{
			subscribe := v1.Group("/subscribe")
			{
				subscribe.POST("/", h.createSubscription)
				subscribe.GET("/:id", h.getSubscription)
				subscribe.PUT("/:id", h.updateSubscription)
				subscribe.DELETE("/:id", h.deleteSubscription)
				subscribe.GET("/all", h.getAllSubscriptions)
				subscribe.GET("/sum", h.getSubsSum)
			}
		}
	}

	return router
}
