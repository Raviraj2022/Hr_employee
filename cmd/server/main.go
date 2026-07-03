package main

import (
	"log"
	"os"

	"example.com/hr-emp-mgmt/config"
	"example.com/hr-emp-mgmt/internal/auth"

	"example.com/hr-emp-mgmt/internal/user"

	"github.com/gin-gonic/gin"
)


func main() {

	config.LoadEnv()

	config.ConnectDatabase()

	router := gin.Default()

	port := os.Getenv("SERVER_PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)
    
	repo := user.NewRepository()
    service := user.NewService(repo)
    handler := auth.NewHandler(service)

    auth.RegisterRoutes(router, handler)
	router.Run(":" + port)
}