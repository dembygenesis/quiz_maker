package seed

import (
	"github.com/dembygenesis/quiz_maker/src/app/common/utilities/string_utils"
	"github.com/dembygenesis/quiz_maker/src/v3/api/modelsV2"
)

func GetUsers() *[]modelsV2.User {
	var users []modelsV2.User
	passwordAdmin, _ := string_utils.Encrypt("admin")
	studentAdmin, _ := string_utils.Encrypt("student")
	users = append(users, modelsV2.User{
		UserTypeID: 1,
		Email:      "coralim@yahoo.com",
		FirstName:  "Cora",
		LastName:   "Lim",
		Gender:     "F",
		Password:   passwordAdmin,
		Address:    "Tubod",
	})
	users = append(users, modelsV2.User{
		UserTypeID: 1,
		Email:      "student@gmail.com",
		FirstName:  "Student First Name",
		LastName:   "Student Last Name",
		Gender:     "M",
		Password:   studentAdmin,
		Address:    "Tubod",
	})
	return &users
}
