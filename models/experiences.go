package models

type Experience struct {
	ID    uint `gorm:"primary_key;auto_increment;not_null"`
	Years uint `gorm:"unique:true"`
}

