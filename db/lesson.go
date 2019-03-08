package db

import _ "github.com/jinzhu/gorm"

type Lesson struct {
	ID    uint   `gorm:"column:id;primary_key" json:"id"`
	Title string `gorm:"column:title" json:"title"`
	Text  string `gorm:"column:text" json:"text"`
}
