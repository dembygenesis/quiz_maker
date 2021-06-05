package migrate

import (
	"fmt"
	"github.com/dembygenesis/quiz_maker/internal/models"
	"gorm.io/gorm"
	"os"
)

func Run(tx *gorm.DB) error {
	fmt.Println("----------- Migrate -----------")
	var err error
	err = tx.Exec("DROP DATABASE " + os.Getenv("DB_DATABASE") + ";").Error
	if err != nil {
		return err
	}
	err = tx.Exec("CREATE DATABASE " + os.Getenv("DB_DATABASE") + ";").Error
	if err != nil {
		return err
	}
	err = tx.Exec("USE " + os.Getenv("DB_DATABASE") + ";").Error
	if err != nil {
		return err
	}
	err =  tx.AutoMigrate(
		&models.UserType{},
		&models.User{},
		&models.Quiz{},
		&models.QuizQuestion{},
		&models.QuizChoice{},
	)
	if err != nil {
		return err
	}
	return nil
}
