package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Bangkok",
		viper.GetString("db.host"),
		viper.GetString("db.user"),
		viper.GetString("db.password"),
		viper.GetString("db.dbname"),
		viper.GetInt("db.port"),
	)

	dbconnect, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	db = dbconnect
}
