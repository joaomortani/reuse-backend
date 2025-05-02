package main

import (
	"fmt"
	"reuse-api/config"
	"reuse-api/database"
	itemHandlers "reuse-api/handlers/item"
	handlers "reuse-api/handlers/user"
	itemRepositories "reuse-api/repositories/item"
	repositories "reuse-api/repositories/user"
	v1 "reuse-api/routes/v1"
	itemServices "reuse-api/services/item"
	services "reuse-api/services/user"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.Connect()

	r := gin.Default()

	// User setup
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Item setup
	itemRepo := itemRepositories.NewItemRepository(database.DB)
	itemService := itemServices.NewItemService(itemRepo)
	itemHandler := itemHandlers.NewItemHandler(itemService)

	v1.RegisterUserRoutes(r, userHandler)
	v1.RegisterItemRoutes(r, itemHandler)

	host := config.GetEnv("APP_HOST")
	port := config.GetEnv("APP_PORT")
	r.Run(fmt.Sprintf("%s:%s", host, port))
}
