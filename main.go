package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/dghubble/oauth1"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// Kodoku is 書き込み用の構造体
type Kodoku struct {
	ID       int       `json:id`
	UserID   string    `json:userid`
	UserName string    `json:username`
	Text     string    `json:text`
	time     time.Time `json:birthday`
	Tweetid  int64     `json:tweetid`
}

type Enmatyou struct {
	ID          int   `json:id gorm:"AUTO_INCREMENT"`
	Tweetid     int64 `json:tweetid`
	Blockwrited bool  `json:blockwrited`
}

type Search_response struct {
	Text string `json:text`
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

	// 閻魔帳への書き込み
	// curl -X POST -H "Content-Type: application/json" -d '{"Tweetid":123456789,"Blockwrited":false}' localhost:3000/enmatyou/write
	r.POST("/enmatyou/write", func(c *gin.Context) {
		data := Enmatyou{}

		if err := c.BindJSON(&data); err != nil {
			c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
		}
		db.NewRecord(data)
		db.Create(&data)
		if db.NewRecord(data) == false {
			c.JSON(http.StatusOK, data)
		}
	})
	// 閻魔帳を全部読む
	r.GET("/enmatyou/read", func(c *gin.Context) {
		contents := []Enmatyou{}
		db.Find(&contents)
		c.JSON(http.StatusOK, contents)
	})
	// 閻魔帳を個数制限降順で読む
	r.GET("/enmatyou/read/:num", func(c *gin.Context) {
		num := c.Param("num")
		contents := []Enmatyou{}
		db.Order("ID desc").Limit(num).Find(&contents)
		c.JSON(http.StatusOK, contents)
	})
	// 閻魔帳のBlockwritedをtrueにする
	// curl -X PUT -H "Content-Type: application/json" -d '{"Blockwrited":true}' localhost:3000/enmatyou/update/1
	r.PUT("/enmatyou/update/:id", func(c *gin.Context) {
		user := Enmatyou{}
		id := c.Param("id")

		data := Enmatyou{}
		if err := c.BindJSON(&data); err != nil {
			c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
		}

		db.Where("ID = ?", id).First(&user).Updates(&data)
	})

	// 蠱毒に書き込み
	r.POST("/kodoku/write", func(c *gin.Context) {
		data := Kodoku{}

		if err := c.BindJSON(&data); err != nil {
			c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
		}
		db.NewRecord(data)
		db.Create(&data)
		if db.NewRecord(data) == false {
			c.JSON(http.StatusOK, data)
		}
	})
	// 蠱毒全部読み込む
	r.GET("/kodoku/read", func(c *gin.Context) {
		contents := []Kodoku{}
		db.Find(&contents)
		c.JSON(http.StatusOK, contents)
	})
	//　蠱毒個数制限で降順に読み込み
	r.GET("/kodoku/read/:num", func(c *gin.Context) {
		num := c.Param("num")
		contents := []Kodoku{}
		db.Order("ID desc").Limit(num).Find(&contents)
		c.JSON(http.StatusOK, contents)
	})

	return r
}

func gettweet() {
	config := oauth1.NewConfig(os.Getenv("TWITTERCONSUMER_KEY"), os.Getenv("TWITTERCONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("TWITTERACCESS_TOKEN"), os.Getenv("TWITTERACCESS_TOKEN_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)

	// fmt.Println(os.Getenv("TWITTERCONSUMER_KEY"))

	request, err := http.NewRequest("GET", "https://api.twitter.com/1.1/statuses/show.json?id=1163100797984366592", nil)
	if err != nil {
		panic(err.Error())
	}

	response, err := httpClient.Do(request)
	if err != nil {
		panic(err.Error())
	}

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}

	var result Search_response
	json.Unmarshal(b, &result)
	response.Body.Close()

	fmt.Println(result.Text)

	// fmt.Println(string(b))
	// fmt.Println(response.Body)
	// fmt.Println(reflect.TypeOf(response.Body))

	// fmt.Println(b[2])

	// fmt.Println(reflect.TypeOf(b))

}

func main() {
	// db := gormConnect()
	// defer db.Close()

	// // 初回マイグレーションで使った
	// db.CreateTable(&Kodoku{})
	// db.CreateTable(&Enmatyou{})

	// port := os.Args[1]
	// r := setRouter(db)
	// r.Run(":" + port)

	gettweet()

}
