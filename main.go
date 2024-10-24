package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jvpr123/golang-crud/src/controller/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":9000"); err != nil {
		log.Fatal(err)
	}
}
