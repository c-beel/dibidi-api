package db

import (
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Lesson{})
	db.AutoMigrate(&LessonDependency{})
}
