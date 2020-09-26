package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Firstname string `gorm:"not null" json:"firstname"`
	Lastname  string `gorm:"not null" json:"lastname"`
	Email     string `gorm:"unique_index;not null" json:"email"`
	Password  string `gorm:"non null" json:"password"`
}

func (u *User) HashPassword() error {
	b := []byte(u.Password)
	hash, err := bcrypt.GenerateFromPassword(b, bcrypt.MinCost)
	if err == nil {
		u.Password = string(hash)
	}
	return err
}
