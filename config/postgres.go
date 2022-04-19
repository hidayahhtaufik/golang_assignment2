package config

import (
	"assignment2/models"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	host   string = "localhost"
	port   int    = 5432
	user   string = "postgres"
	pass   string = "postgres"
	dbname string = "gosql"
)

func NewPostgres() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	db.Debug().AutoMigrate(models.Order{}, models.Item{})

	return db
}

func GetDB() *gorm.DB {
	var db *gorm.DB
	return db
}
