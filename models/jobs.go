package models

import "time"

type Job struct {
	ID                  uint   `gorm:"primaryKey;autoIncrement:true"`
	Title               string `gorm:"typevarchar(100);unique"`
	Description         string `gorm:"typevarchar(500)"`
	PublisherReferer    uint
	Publisher           Publisher `gorm:"foreignKey:PublisherReferer"`
	Link                string    `gorm:"typevarchar(500)"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	PublishedAt         time.Time
	FinishedAt          time.Time
	CompanyReferer      uint
	Company             Company `gorm:"foreignKey:CompanyReferer"`
	EnglishLevelReferer uint
	EnglishLevel        EnglishLevel `gorm:"foreignKey:EnglishLevelReferer"`
	Experience          uint
	SchedulesReferer    uint
	Schedules           Schedules          `gorm:"foreignKey:SchedulesReferer"`
	ProgrammingSkills   []ProgrammingSkill `gorm:"many2many:job_programming_skills;"`
	PersonalSkills      []PersonalSkill    `gorm:"many2many:job_personal_skills;"`
	Salary              uint
}
