package helpers

import (
	"backend/models"
	"fmt"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"crypto/sha256"
	"encoding/hex"
)

func OpenDb() (*gorm.DB, error) {
	db_url := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(db_url), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func MakeMigrations() {
	db, err := OpenDb()
	if err != nil {fmt.Println("FAILED TO MIGRATE WITH ERROR:", err.Error())}
	go db.AutoMigrate(&models.User{}) //Paralellism go brrr
	go db.AutoMigrate(&models.Like{})
	go db.AutoMigrate(&models.Post{})
	fmt.Println("MIGRATED THE DATABASE")
}

func Sha256(entry string) string {
	result := sha256.Sum256([]byte(entry))
	return hex.EncodeToString(result[:])
}