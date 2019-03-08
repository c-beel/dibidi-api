package db

import _ "github.com/jinzhu/gorm"

type LessonDependency struct {
	ID          uint   `gorm:"column:id;primary_key" json:"id"`
	Source      Lesson `gorm:"column:source" json:"source"`
	Destination Lesson `gorm:"column:destination" json:"destination"`
}
