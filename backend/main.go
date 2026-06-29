package main

import (
	"kampus-api/config"
	"kampus-api/routes"

	_ "kampus-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/cors"
)

// @title Kampus API
// @version 1.0
// @description API CRUD Mahasiswa dan Jurusan
// @host localhost:8080
// @BasePath /

func main() {

	config.ConnectDB()

	r := gin.Default()

	r.Use(cors.Default())

	routes.SetupRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
