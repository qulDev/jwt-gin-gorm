package migration

import (
	"log"

	"github.com/qulDev/jwt-gin-gorm/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}
}
