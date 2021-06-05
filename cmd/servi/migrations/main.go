package main

import (
	"errors"
	"fmt"
	"github.com/dembygenesis/quiz_maker/cmd/servi/migrations/migrate"
	"github.com/dembygenesis/quiz_maker/cmd/servi/migrations/seed"
	"github.com/dembygenesis/quiz_maker/internal/platform/mysql"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
)

func init() {
	err := godotenv.Load()
	fmt.Println("============Loaded============")

	if err != nil {
		panic("Error loading .env files:" + err.Error())
	}
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var err error
	db, err := mysql.GetGormInstance()
	if err != nil {
		return errors.New("error fetching gorm instance: " + err.Error())
	}
	err = db.Transaction(func(tx *gorm.DB) error {
		err = migrate.Run(tx)
		if err != nil {
			return err
		}
		err = seed.Run(tx)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}