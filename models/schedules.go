package models

type Schedules struct {
	ID   uint   `gorm:"primaryKey;autoIncrement:true"`
	Name string `gorm:"typevarchar(10);unique:true"`
}