package migration

import (
	"example.com/awesome/model"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func AddUserTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20240622_add_user_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.Migrator().CreateTable(&model.User{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&model.User{})
		},
	}
}
