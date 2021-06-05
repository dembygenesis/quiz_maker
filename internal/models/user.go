package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserTypeID int
	FirstName  string
	LastName   string
	Gender     string
	UserType   UserType
}
