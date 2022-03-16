package models

type Publisher struct {
	ID   uint   `gorm:"primaryKey;autoIncrement:true"`
	Name string `gorm:"typevarchar(20);unique:true"`
}
