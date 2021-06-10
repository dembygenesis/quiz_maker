package modelsV2

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserTypeID int
	FirstName  string
	Email      string
	LastName   string
	Password   string
	Gender     string
	Address    string
	UserType   UserType
}
