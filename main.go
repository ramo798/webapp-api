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

	// 全部読み込む
	r.GET("/readcontents", func(c *gin.Context) {
		contents := []Zigokucontents{}
		db.Find(&contents)
		c.JSON(http.StatusOK, contents)
	})
	//　個数制限で降順に読み込み
	r.GET("/readcontents/:num", func(c *gin.Context) {
		num := c.Param("num")
		contents := []Zigokucontents{}
		db.Order("ID desc").Limit(num).Find(&contents)
		c.JSON(http.StatusOK, contents)
	})

	// postエンドポイント
	r.POST("/writecontents", func(c *gin.Context) {
		data := Zigokucontents{}

		if err := c.BindJSON(&data); err != nil {
			c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
		}
		db.NewRecord(data)
		db.Create(&data)
		if db.NewRecord(data) == false {
			c.JSON(http.StatusOK, data)
		}
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
