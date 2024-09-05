package database

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       uint   `gorm:"primaryKey;auto_increment" json:"id"`
	Name     string `gorm:"size:30" json:"name"`
	Email    string `gorm:"size:30;unique" json:"email"`
	Password string `gorm:"size:255" json:"password"`
}

func (u *User) SetPassword() {
	// password -> hash
	pwHash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	// save
	u.Password = string(pwHash)
}
