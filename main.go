package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// Zigokucontents is 書き込み用の構造体
type Zigokucontents struct {
	ID       int    `json:id gorm:"AUTO_INCREMENT"`
	UserID   string `json:userid`
	UserName string `json:username`
	Text     string `json:text`
	Year     int    `json:year`
	Month    int    `json:year`
	Day      int    `json:year`
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

func setRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/testread", func(c *gin.Context) {
		users := []Zigokucontents{}
		db.Find(&users)
		c.JSON(http.StatusOK, users)
	})

	r.GET("/testwrite", func(c *gin.Context) {

	})

	return r
}

func main() {
	db := gormConnect()
	defer db.Close()

	// 初回マイグレーションで使った
	// db.CreateTable(&Zigokucontents{})

	port := os.Args[1]
	r := setRouter(db)
	r.Run(":" + port)

}
