package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string `gorm:"uniqueIndex"`
	PasswordHashed string `gorm:"type:varchar(255) or null"`
}

func (User) TableName() string {
	return "user"
}