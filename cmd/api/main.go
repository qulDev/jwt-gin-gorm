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

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)

	// Initialize Service
	userService := service.NewUserService(userRepo)

	// Initialize Handlers
	userHandler := handler.NewUserHandler(userService)
	articleHandler := handler.NewArticleHandler()

	// Initialize Gin router
	r := routes.SetupRoutes(&routes.Handlers{
		User:    userHandler,
		Article: articleHandler,
	})

	r.Run(":8080")
}
