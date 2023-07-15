package routes

import (
	"svi-backend/controllers"

	"github.com/gin-gonic/gin"
)

func CreatePostRoutes(c *gin.Engine, controllers controllers.PostControllerInterface) {
	ur := c.Group("/api/v1/articles")
	ur.GET("/", controllers.Get)
	ur.GET("/:id", controllers.GetById)
	ur.POST("/", controllers.Create)
	ur.PUT("/:id", controllers.Update)
	ur.DELETE("/:id", controllers.Delete)
}
