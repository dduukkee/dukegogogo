package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"io/ioutil"
	"net/http"
)

type Product struct {
	gorm.Model
	Code    string
	weather uint
}

func main() {
	//////////db	//////////
	db, err := gorm.Open("sqlite3", "test.db")
	// 	if err != nil {
	// 		panic("failed to connect database")
	// 	}
	// 	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// 创建
	db.Create(&Product{Code: "api", weather: 1000})

	// 读取
	var product Product
	db.First(&product, "code = ?", "api") // 查询code为api的product
	ttt := db.Where("code = ?", "api").First(&product)
	fmt.Printf("Get weather: %v \n", ttt)

	// 更新 - 更新product的price为2000
	db.Model(&product).Update("weather", 2000)

	//api
	resp, _ := http.Get("http://weather.json.tw/api")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	// 	fmt.Println(string(body))

	//////////redis	//////////
	c, err := redis.Dial("tcp", "127.0.0.1:6379")

	_, err = c.Do("SET", "weather", string(body))
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	// 	username, err := redis.String(c.Do("GET", "weather"))
	// 	if err != nil {
	// 		fmt.Println("redis get failed:", err)
	// 	} else {
	// 		fmt.Printf("Get weather: %v \n", username)
	// 	}

}
