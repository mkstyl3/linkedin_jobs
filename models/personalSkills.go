package models 

type PersonalSkill struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:true"`
	Name        string `gorm:"typevarchar(30);unique:true"`
}