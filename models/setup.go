package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Database string
	Password string
	SSL      string
}

func ConnectDatabase() {

	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Database: os.Getenv("DB_DATABASE"),
		Password: os.Getenv("DB_PASSWORD"),
		SSL:      os.Getenv("DB_SSL"),
	}
	dbURL := fmt.Sprintf(

		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Database,
		dbConfig.Password,
		dbConfig.SSL,
	)
	fmt.Println(dbURL)

	database, err := gorm.Open("postgres", dbURL)

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Book{}, &User{})

	DB = database
}
