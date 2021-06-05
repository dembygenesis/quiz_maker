package models

type QuizChoice struct {
	ID             int
	QuizQuestionID int `gorm:"uniqueIndex:uniq_composite;type:varchar(255);"`
	QuizQuestion   QuizQuestion
	Name           string `gorm:"uniqueIndex:uniq_composite;type:varchar(255);"`
	Slug           string `gorm:"uniqueIndex:uniq_composite;type:varchar(255);"`
	Order          int    `gorm:"uniqueIndex:uniq_composite;type:varchar(255);"`
}
