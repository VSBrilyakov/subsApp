package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		subscribe := api.Group("/subscribe")
		{
			subscribe.POST("/", h.createSubscribe)
			subscribe.GET("/:id", h.getSubscribe)
			subscribe.PUT("/:id", h.updateSubscribe)
			subscribe.DELETE("/:id", h.deleteSubscribe)
			subscribe.GET("/all", h.getAllSubscribes)
		}
	}

	return router
}
