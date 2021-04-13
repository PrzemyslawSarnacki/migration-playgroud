package main

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var initialMigration = &gormigrate.Migration{
	ID: "1257894000",
	Migrate: func(tx *gorm.DB) error {
		err := tx.AutoMigrate(&User{})
		return err
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("users")
	},
}