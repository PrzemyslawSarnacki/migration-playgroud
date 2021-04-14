package migrations

import (
	"fmt"
	"time"

	"https://github.com/PrzemyslawSarnacki/migration-playgroud/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

var timeStamp = fmt.Sprintf("%d", time.Now().Unix())

var timestampMigration = &gormigrate.Migration{
	ID: timeStamp,
	Migrate: func(tx *gorm.DB) error {
		err := tx.AutoMigrate(&models.User{})
		tx.Create(&models.User{
			Name: "Ken",
			Data: datatypes.JSON([]byte(`{"foo1": "bar1", "foo2": "bar3"}`)),
		})
		tx.Create(&models.User{
			Name: "Jack",
			Data: datatypes.JSON([]byte(`{"foo1": "bar1", "foo2": "bar3"}`)),
		})
		return err
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("users")
	},
}
