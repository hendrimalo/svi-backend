package main

import (
	"svi-backend/configs"
	"svi-backend/controllers"
	"svi-backend/repository"
	"svi-backend/routes"
	"svi-backend/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	//articles routes & depedency injection
	db := configs.OpenMySQL()
	pRepo := repository.NewPostRepository(db)
	pService := services.NewPostService(db, pRepo)
	pController := controllers.NewPostController(pService)
	routes.CreatePostRoutes(r, pController)

	r.Run("0.0.0.0:3000")
}
