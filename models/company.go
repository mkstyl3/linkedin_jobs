package models

type CompanySize struct {
	ID   uint   `gorm:"primaryKey;autoIncrement:true"`
	Name string `gorm:"typevarchar(10);unique:true"`
}

type Company struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:true"`
	Name        string `gorm:"typevarchar(20);unique:true"`
	SizeReferer uint
	Size        CompanySize `gorm:"foreignKey:SizeReferer"`
}

// func GetCompanySizes() ([]CompanySize, error) {
// 	companySizes := []CompanySize{}
// 	result := Db.Find(&companySizes)
// 	return companySizes, result.Error
// }


