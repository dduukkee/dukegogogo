// 題目：
// 撰寫一程式可對目標API 進行連線，需可自行調整併發數量

// 記錄：
// 單一連線花費時間
// 單一最大花費時間
// 單一最小花費時間
// 壓測總數花費時間
// 平均花費時間

// cron, slice, (t Time) Sub, goroutine

package main

import (
	"fmt"
	// "github.com/robfig/cron"
	"io/ioutil"
	// "log"
	"net/http"
)

func main() {
	api()
	// i := 0
	// c := cron.New()
	// spec := "*/1 * * * * *"
	// c.AddFunc(spec, func() {
	// 	i++
	// 	log.Println("execute per second", i)
	// })
	// c.Start()
	// select {}
}

func api() {
	resp, _ := http.Get("http://bb.dev.d2/api/domain/32/levels")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
