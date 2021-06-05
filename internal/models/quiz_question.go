package models

type QuizQuestion struct {
	ID     uint
	QuizId int
	Quiz   Quiz
	Name   string `gorm:"uniqueIndex:uniq_name;type:varchar(255);"`
	Slug   string `gorm:"uniqueIndex:uniq_slug;type:varchar(255);"`
	Answer string `gorm:"uniqueIndex:uniq_slug;type:varchar(255);"`
	Order  int    `gorm:"uniqueIndex:uniq_order;type:varchar(255);"`
}
