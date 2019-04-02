package main

import (
	"github.com/c-beel/dibidi-api/db"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/c-beel/dibidi-api/server"
	"github.com/gin-contrib/cors"
	_ "time"
)

func main() {
	var err error
	var app server.Application
	app.DB, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Println("Can't connect to database. :(")
		fmt.Println(err)
		os.Exit(1)
	}
	defer app.DB.Close()
	db.Migrate(app.DB)

	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}))
	engine.POST("/lessons/add", app.AddLesson)
	engine.POST("/lessons/edit", app.EditLesson)
	engine.GET("/lessons/", app.GetLessonsList)
	engine.GET("/lesson/", app.GetLesson)
	engine.DELETE("/lessons/delete", app.DeleteLesson)

	engine.Run(":8080")
}
