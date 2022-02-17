package models

type Comment struct {
	ID   uint   `gorm:"primary_key;auto_increment;not_null"`
	Text string `gorm:"typevarchar(100)"`
}

func GetCommentsDb() ([]Comment, error) {
	var comments []Comment
	if err := Db.Find(&comments).Error; err != nil {
		return comments, err
	}
	return comments, nil
}

func PostCommentsDb(comment Comment) (Comment, error) {
	if err := Db.Create(&comment).Error; err != nil {
		return comment, err
	}
	return comment, nil
}
