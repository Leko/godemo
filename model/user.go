package model

import (
	"github.com/jinzhu/gorm"
	"godemo/database"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model

	Email    string `sql:"not null;unique_index"`
	Password string `sql:"not null"`
}

func PasswordHash(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashed)
}

func (u *User) Auth() (int, error) {
	email := u.Email
	password := u.Password

	db := database.GetDB()
	user := User{}

	db.Where(&User{Email: email}).Find(&user)

	// https://godoc.org/golang.org/x/crypto/bcrypt#CompareHashAndPassword
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return int(user.ID), err
}
