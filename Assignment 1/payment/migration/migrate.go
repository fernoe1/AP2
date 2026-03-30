package migration

import (
	"github.com/fernoe1/AP2/assignment-1/payment/internal/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&domain.Payment{})
}
