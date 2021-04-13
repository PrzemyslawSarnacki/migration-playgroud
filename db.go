package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"

	"gorm.io/datatypes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	var err error

	if DB, err = OpenTestConnection(); err != nil {
		log.Printf("failed to connect database, got error %v\n", err)
		os.Exit(1)
	} else {
		sqlDB, err := DB.DB()
		if err == nil {
			err = sqlDB.Ping()
		}

		if err != nil {
			log.Printf("failed to connect database, got error %v\n", err)
		}

		RunMigrations()

		if DB.Dialector.Name() == "sqlite" {
			DB.Exec("PRAGMA foreign_keys = ON")
		}

		DB.Logger = DB.Logger.LogMode(logger.Info)
	}

}

func OpenTestConnection() (db *gorm.DB, err error) {
	dbDSN := os.Getenv("GORM_DSN")
	log.Println("testing postgres...")
	if dbDSN == "" {
		dbDSN = "user=gorm password=gorm host=localhost dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	}
	db, err = gorm.Open(postgres.Open(dbDSN), &gorm.Config{})

	if debug := os.Getenv("DEBUG"); debug == "true" {
		db.Logger = db.Logger.LogMode(logger.Info)
	} else if debug == "false" {
		db.Logger = db.Logger.LogMode(logger.Silent)
	}

	return
}

func RunMigrations() {
	var err error
	allModels := []interface{}{&User{}}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(allModels), func(i, j int) { allModels[i], allModels[j] = allModels[j], allModels[i] })

	var migrations = []*gormigrate.Migration{
		{
			ID: "1257894000",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("users")
			},
		},
	}
	timeStamp := fmt.Sprintf("%d", time.Now().Unix())

	extendedMigrations := append(migrations, &gormigrate.Migration{
		ID: timeStamp,
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&User{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("users")
		},
	})

	m := gormigrate.New(DB, gormigrate.DefaultOptions, extendedMigrations)

	if err = m.RollbackLast(); err != nil {
		log.Fatalf("Could not Rollback: %v", err)
	}

	if err = m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	for _, m := range allModels {
		if !DB.Migrator().HasTable(m) {
			log.Printf("Failed to create table for %#v\n", m)
			os.Exit(1)
		}
	}
}

func GetMultipleJSONQuery(queries ...*datatypes.JSONQueryExpression) ([]User, error) {
	var users []User

	if err := DB.Where(queries).Find(&users).Error; err != nil {
		return nil, errors.New("failed to execute json query")
	}

	return users, nil
}
