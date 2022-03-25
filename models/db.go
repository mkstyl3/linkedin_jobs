package models

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func ConnectToDb() error {
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"))
	Db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	return err
}

func InitDB() error {
	return ConnectToDb()
}

func GetAll(models interface{}) error {
	return Db.Find(models).Error
}

func GetByAttr(m interface{}, attrName string, attrValue string) error {
	return Db.Where(attrName+" = ?", attrValue).First(m).Error
}

func Insert(model interface{}) error {
	return Db.Create(model).Error
}
