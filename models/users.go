package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidLogin = errors.New("invalid login")
)

type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement:true"`
	Name      string `gorm:"typevarchar(20);unique:true"`
	Password  []byte `gorm:"size:32"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func AuthenticateUser(username, password string) error {
	var user User
	user.Name = username
	user, err := GetUser(user)
	if err != nil {
		return ErrUserNotFound
	}
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		return ErrInvalidLogin
	} else {
		return err
	}
}

func GetUser(user User) (User, error) {
	if err := Db.Where("name = ?", user.Name).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Info().Err(err).Msg(fmt.Sprintf("User %s not found", user.Name))
			return user, ErrUserNotFound
		} else {
			log.Fatal().Err(err).Msg(fmt.Sprintf("User %s not found", user.Name))
		}
	}
	fmt.Printf("Username retrieved: %s", user.Name)
	return user, err
}
