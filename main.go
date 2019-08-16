package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Args[1]
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	r.Run(":" + port)
}

// package main

// import (
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/ant0ine/go-json-rest/rest"
// 	_ "github.com/go-sql-driver/mysql" // エイリアスでprefixを省略できる
// 	"github.com/jinzhu/gorm"
// )

// type Country struct {
// 	Id        int64     `json:"id"`
// 	Name      string    `sql:"size:1024" json:"name"`
// 	CreatedAt time.Time `json:"createdAt"`
// }

// type Impl struct {
// 	DB *gorm.DB
// }

// func (i *Impl) InitDB() {
// 	var err error
// 	// MySQLとの接続。ユーザ名：gorm パスワード：password DB名：country
// 	i.DB, err = gorm.Open("mysql", "gorm:password@/country?charset=utf8&parseTime=True&loc=Local")
// 	if err != nil {
// 		log.Fatalf("Got error when connect database, the error is '%v'", err)
// 	}
// 	i.DB.LogMode(true)
// }

// // DBマイグレーション
// func (i *Impl) InitSchema() {
// 	i.DB.AutoMigrate(&Country{})
// }

// func main() {

// 	i := Impl{}
// 	i.InitDB()
// 	i.InitSchema()

// 	api := rest.NewApi()
// 	api.Use(rest.DefaultDevStack...)
// 	router, err := rest.MakeRouter()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Printf("server started.")
// 	api.SetApp(router)
// 	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
// }
