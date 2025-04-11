package main

import (
	"fmt"
	"reuse-api/config"
	"reuse-api/database"
	handlers "reuse-api/handlers/user"
	repositories "reuse-api/repositories/user"
	v1 "reuse-api/routes/v1"
	services "reuse-api/services/user"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.Connect()

	r := gin.Default()

	repo := repositories.NewUserRepository()
	service := services.NewUserService(repo)
	handler := handlers.NewUserHandler(service)

	v1.RegisterUserRoutes(r, handler)

	host := config.GetEnv("APP_HOST")
	port := config.GetEnv("APP_PORT")
	r.Run(fmt.Sprintf("%s:%s", host, port))
}
