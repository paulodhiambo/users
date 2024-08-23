package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"size:100;not null"`
	LastName  string `gorm:"size:100;not null"`
	Email     string `gorm:"size:100;uniqueIndex;not null"`
	Gender    string `gorm:"size:10;not null"`
	IPAddress string `gorm:"size:50;not null"`
	Password  string `gorm:"size:255;not null"`
}
