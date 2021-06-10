package seed

import (
	"fmt"
	"github.com/dembygenesis/quiz_maker/src/app/common/utilities/db"
	"github.com/dembygenesis/quiz_maker/src/v3/api/modelsV2"
	"gorm.io/gorm"
)

<<<<<<< HEAD
func ProcessUsers(tx *gorm.DB, u *[]modelsV2.User) error {
	return tx.Create(u).Error
}

func ProcessUserTypes(tx *gorm.DB, u []string) error {
	var userTypes []modelsV2.UserType
	for _, v := range u {
		userTypes = append(userTypes, modelsV2.UserType{
			Name: v,
		})
	}
	return tx.Create(userTypes).Error
}

=======
>>>>>>> e44fe428e2750bbc1edf62da92fd337610f90487
func ProcessQuiz(tx *gorm.DB, q *Quiz) error {
	var err error

	quiz := modelsV2.Quiz{
		Model:    gorm.Model{},
		Name:     q.Name,
		Slug:     q.Slug,
		Order:    q.Order,
		Duration: q.Duration,
	}
	// Insert quiz
	err = tx.Create(&quiz).Error
	if err != nil {
		return err
	}

	lastInsertQuizId, err := db.GetLastInsertIDGorm(tx)
	if err != nil {
		return err
	}
	for quizQuestionOrder, v := range q.QuizQuestion {
		quizQuestionOrder++
		quizQuestion := QuizQuestion{
			QuizID: lastInsertQuizId,
			Name:   v.Name,
			Slug:   v.Slug,
			Answer: v.Answer,
			Order:  quizQuestionOrder,
		}
		// Insert quiz question
		err = tx.Omit("quiz_choice").Create(&quizQuestion).Error
		if err != nil {
			return err
		}
		lastInsertQuizQuestionId, err := db.GetLastInsertIDGorm(tx)
		if err != nil {
			return err
		}
		for quizChoiceOrder, y := range v.QuizChoice {
			quizChoiceOrder++
			quizChoice := QuizChoice{
				QuizQuestionID: lastInsertQuizQuestionId,
				Name:           y.Name,
				Slug:           y.Slug,
				Order:          quizChoiceOrder,
			}
			// Insert quiz choice
			err = tx.Create(&quizChoice).Error
			if err != nil {
				return err
			}
		}
	}
	return nil
}

<<<<<<< HEAD
// go build -i -o C:\Users\paolo\go\custom_output_dir

=======
>>>>>>> e44fe428e2750bbc1edf62da92fd337610f90487
func Run(tx *gorm.DB) error {
	fmt.Println("----------- Seeding -----------")

	var err error
	quiz := GetDataQuiz1()
<<<<<<< HEAD
	userTypes := GetUserTypes()
	users := GetUsers()
=======
>>>>>>> e44fe428e2750bbc1edf62da92fd337610f90487

	err = ProcessQuiz(tx, &quiz)
	if err != nil {
		return err
	}
<<<<<<< HEAD

	err = ProcessUserTypes(tx, userTypes)
	if err != nil {
		return err
	}

	err = ProcessUsers(tx, users)
	if err != nil {
		return err
	}

=======
>>>>>>> e44fe428e2750bbc1edf62da92fd337610f90487
	return nil
}
