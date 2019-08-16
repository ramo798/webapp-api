package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// WriteContents is 書き込み用の構造体
type WriteContents struct {
	ID       int    `json:id sql:AUTO_INCREMENT`
	UserID   string `json:userid`
	UserName string `json:username`
	Date     string `json:date`
}

func gormConnect() *gorm.DB {
	url := os.Getenv("DATABASE_URL")

	connection, err := pq.ParseURL(url)
	if err != nil {
		panic(err.Error())
	}
	connection += " sslmode=require"

	db, err := gorm.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("db connected: ", &db)
	return db
}

func main() {
	db := gormConnect()
	defer db.Close()
	// i := Impl{}
	// i.InitDB()

	// マイグレーション
	// db.CreateTable(&WriteContents{})

	// insert test
	// cont := WriteContents{}
	// cont.Date = "20190816"
	// cont.UserID = "asasa"
	// cont.UserName = "name"
	// db.Create(&cont)

	// port := os.Args[1]
	// r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "Hello World")
	// })

	// r.Run(":" + port)
}
