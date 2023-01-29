package main

import (
	"fmt"
	"hahu_jobs_projects/routes"
	"log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serveApplication() {
	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/images", "./images")
	routes.UseRoute(router)
	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
