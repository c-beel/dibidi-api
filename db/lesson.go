package db

import _ "github.com/jinzhu/gorm"

type LessonListItem struct {
	ID    uint   `gorm:"column:id;primary_key" json:"id"`
	Title string `gorm:"column:title" json:"title"`
}

type Lesson struct {
	LessonListItem
	Text string `gorm:"column:text" json:"text"`
}

func (Lesson) TableName() string {
	return "lessons"
}
