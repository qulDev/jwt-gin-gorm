package main

import (
	"github.com/qulDev/jwt-gin-gorm/internal/config"
	"github.com/qulDev/jwt-gin-gorm/internal/handler"
	"github.com/qulDev/jwt-gin-gorm/internal/migration"
	"github.com/qulDev/jwt-gin-gorm/internal/repository"
	"github.com/qulDev/jwt-gin-gorm/internal/routes"
	"github.com/qulDev/jwt-gin-gorm/internal/service"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize database and cache
	db := config.InitDB()
	cache := config.InitCache()

	// Run database migrations
	migration.Migrate(db)

	_ = cache // to avoid unused variable error, remove when cache is used

	// Initialize repositories, services, and handlers
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHendler := handler.NewUserHandler(userService)
	articleHendler := handler.NewArticleHandler()

	// Initialize Gin router
	r := routes.SetupRoutes(&routes.Handlers{
		User:    userHendler,
		Article: articleHendler,
	})

	r.Run(":8080")
}
