package handler

import "github.com/gin-gonic/gin"

type handler struct {
}

func (h *handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up")
		auth.POST("/sign-in")
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/")
			lists.GET("/")
			lists.GET("/:id")
			lists.PUT("/:id")
			lists.DELETE("/:id")

			items := lists.Group("/:id/items")
			{
				items.POST("/")
				items.GET("/")
				items.GET("/:id")
				items.PUT("/:id")
				items.DELETE("/:id")

			}
		}
	}

	return router
}
