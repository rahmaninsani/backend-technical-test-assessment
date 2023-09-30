package app

import (
	"fmt"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/config"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewDB() *gorm.DB {
	constant := config.Constant
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		constant.DBHost, constant.DBUser, constant.DBPassword, constant.DBName, constant.DBPort, constant.DBSSLMode, constant.DBTimezone)
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database\n", err.Error())
	}
	
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	
	models := []interface{}{
		&domain.User{},
		&domain.Post{},
		&domain.Category{},
		&domain.Tag{},
		&domain.PostCategory{},
		&domain.PostTag{},
	}
	
	err = db.AutoMigrate(models...)
	if err != nil {
		log.Fatal("Migration Failed:\n", err.Error())
	}
	
	log.Println("🚀 Connected successfully to the database")
	
	return db
}
