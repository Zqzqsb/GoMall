package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID             uint   `gorm:"primaryKey"`
	Email          string `gorm:"uniqueIndex;type:varchar(255);not null"`
	PasswordHashed string `gorm:"type:varchar(255);null"`
}

func (User) TableName() string {
	return "user"
}

func Create(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

func GetbyEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	return &user, err
}