package mysql

import (
	"fmt"
	"github.com/dembygenesis/quiz_maker/internal/platform/utils/string_utils"
	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"strings"
)

type DB struct {
	Dialect  string `envconfig:"DIALECT" default:"mysql"`
	Host     string `envconfig:"DB_HOST"`
	Username string `envconfig:"DB_USER"`
	Password string `envconfig:"DB_PASSWORD"`
	Name     string `envconfig:"DB_DATABASE"`
	Charset  string `envconfig:"CHARSET" default:"utf8"`
	Timezone string `envconfig:"TIMEZONE"`
}

var (
	// Handle to DB connection
	Handle *gorm.DB
)

func GetGormInstance() (*gorm.DB, error) {
	var dbConfig DB
	err := envconfig .Process("partner", &dbConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string_utils.Dump(dbConfig))

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=%s&parseTime=True&loc=%s",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Name,
		dbConfig.Charset,
		strings.Replace(dbConfig.Timezone, "/", "%2F", -1))

	// Let's see you implement
	namingStrategy := schema.NamingStrategy{
		SingularTable: false,
	}

	Handle, err = gorm.Open(mysql.Open(dbURI), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: namingStrategy,
	})

	if err != nil {
		log.Fatal("Could not connect database")
	} else {
		log.Println("============== Successfully connected ==============", Handle)
	}

	Handle.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8 auto_increment=1")

	return Handle, nil
}

func GetLastInsertIDGorm(tx *gorm.DB) (int, error) {
	var lastInsertId int
	err := tx.Raw(`SELECT LAST_INSERT_ID()`).Scan(&lastInsertId).Error
	return lastInsertId, err
}
