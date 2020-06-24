package migration

import (
	"github.com/arifseft/go-actions/database/entity"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(entity.User{})
}
