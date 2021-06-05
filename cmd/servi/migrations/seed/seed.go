package seed

import (
	"fmt"
	"github.com/dembygenesis/quiz_maker/internal/models"
	"github.com/dembygenesis/quiz_maker/internal/platform/mysql"
	"gorm.io/gorm"
)

func ProcessQuiz(tx *gorm.DB, q *Quiz) error {
	var err error

	quiz := models.Quiz{
		Model: gorm.Model{},
		Name:  q.Name,
		Slug:  q.Slug,
		Order: q.Order,
	}
	// Insert quiz
	err = tx.Create(&quiz).Error
	if err != nil {
		return err
	}

	lastInsertQuizId, err := mysql.GetLastInsertIDGorm(tx)
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
		lastInsertQuizQuestionId, err := mysql.GetLastInsertIDGorm(tx)
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

func Run(tx *gorm.DB) error {
	fmt.Println("----------- Seeding -----------")

	var err error
	quiz := GetDataQuiz1()

	err = ProcessQuiz(tx, &quiz)
	if err != nil {
		return err
	}
	return nil
}
