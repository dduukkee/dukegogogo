package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"io/ioutil"
	"net/http"
)

type Weather struct {
	gorm.Model
	Code string
	Data string
}

func main() {

	//////////redis exist//////////
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	redisData, err := redis.String(c.Do("GET", "weather"))

	if err != nil {
		//////////Db exist//////////
		db, _ := gorm.Open("sqlite3", "weather.db")
		db.AutoMigrate(&Weather{})
		var weather []Weather
		db.Where("Code = ?", "weather").Find(&weather)
		if weather == nil {
			//////////api get data//////////
			resp, _ := http.Get("http://weather.json.tw/api")
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(string(body))

			//api write to redis
			_, err = c.Do("SET", "weather", string(body))

			//api write to db
			db.Create(&Weather{Code: "weather", Data: string(body)})
			defer db.Close()

			fmt.Printf("Get weather from api: %v \n", weather)
		} else {
			fmt.Printf("Get weather from db: %v \n", weather)
		}
	} else {
		fmt.Printf("Get weather from redis: %v \n", redisData)
	}

}
