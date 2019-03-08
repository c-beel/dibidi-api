package main

import (
	"github.com/Dibidi/dibidi-api/db"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/Dibidi/dibidi-api/server"
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
			//return origin == "http://localhost:3000"
		},
	}))
	//engine.Use(cors.Default())
	engine.GET("/lessons/", app.GetLessonsList)
	engine.GET("/lesson/", app.GetLesson)

	engine.Run(":8080")
}
