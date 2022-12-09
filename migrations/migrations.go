package migrations

import (
	"github.com/jaquelineabreu/livraria-go-api-rest/models"
	"gorm.io/gorm"
)


func RunMigrations(db *gorm.DB){
	db.AutoMigrate(models.Book{})
}