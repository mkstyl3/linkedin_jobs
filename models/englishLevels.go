package models

type EnglishLevels struct {
	ID    uint   `gorm:"primaryKey;autoIncrement:true"`
	Level string `gorm:"typevarchar(20);unique:true"`
}
