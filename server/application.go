package server

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"github.com/c-beel/dibidi-api/db"
	"fmt"
)

type Application struct {
	DB *gorm.DB
}

func (app *Application) GetLessonsList(c *gin.Context) {
	var lessons []db.LessonListItem
	if err := app.DB.Table(db.Lesson{}.TableName()).Find(&lessons).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, lessons)
	}
}

func (app *Application) GetLesson(c *gin.Context) {
	getItems := c.Request.URL.Query()
	id := getItems.Get("id")
	var lesson db.Lesson
	if err := app.DB.Where("id = ?", id).First(&lesson).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, lesson)
	}
}

func (app *Application) AddLesson(c *gin.Context) {
	var lesson db.Lesson
	c.BindJSON(&lesson)
	lesson.ID = 0
	if err := app.DB.Create(&lesson).Error; err != nil {
		c.AbortWithStatus(500)
		fmt.Println(app.DB.Error)
	} else {
		c.JSON(200, lesson)
	}
}

func (app *Application) EditLesson(c *gin.Context) {
	var lesson db.Lesson
	c.BindJSON(&lesson)
	if err := app.DB.Model(&lesson).Where("id = ?", lesson.ID).Updates(lesson).Error; err != nil {
		c.AbortWithStatus(500)
		fmt.Println(app.DB.Error)
	} else {
		c.JSON(200, lesson)
	}
}
