package database

import (
	"github.com/overlineink/nellcorp-challenge-jorge-costa/domain/entities"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)

	err := godotenv.Load(basePath + "/../../.env")

	if err != nil {
		log.Fatalf("Error loading .env files")
	}
}

func ConnectDB() *gorm.DB {
	dsn := os.Getenv("DSN")

	db, err := gorm.Open(os.Getenv("DB_TYPE"), dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	if os.Getenv("DEBUG") == "true" {
		db.LogMode(true)
	}

	if os.Getenv("AUTO_MIGRATE_DB") == "true" {
		db.AutoMigrate(&entities.Transaction{}, &entities.Account{})
	}

	return db
}
