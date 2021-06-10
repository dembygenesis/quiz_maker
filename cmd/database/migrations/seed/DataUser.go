package seed

import (
	"github.com/dembygenesis/quiz_maker/src/app/common/utilities/string_utils"
	"github.com/dembygenesis/quiz_maker/src/v3/api/modelsV2"
)

func GetUsers() *[]modelsV2.User {
	var users []modelsV2.User
	password, _ := string_utils.Encrypt("admin")
	users = append(users, modelsV2.User{
		UserTypeID: 1,
		Email:      "coralim@yahoo.com",
		FirstName:  "Cora",
		LastName:   "Lim",
		Gender:     "F",
		Password:   password,
		Address:    "Tubod",
	})
	return &users
}
