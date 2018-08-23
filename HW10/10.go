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
	"github.com/robfig/cron"
	// "io/ioutil"
	"net/http"
	"time"
)

func main() {
	round := 0 //回合
	var timeSlice []time.Duration
	var max time.Duration
	var min time.Duration
	c := cron.New()
	spec := "*/1 * * * * *"
	c.AddFunc(spec, func() {
		round++
		start := time.Now()
		go apiRoutine(round)
		t := time.Now()
		elapsed := t.Sub(start)
		timeSlice = append(timeSlice, elapsed)
		if round == 1 {
			min = elapsed
		}
		//all
		fmt.Println("單一連線花費時間", elapsed)

		//max min
		if elapsed > max {
			max = elapsed
		}
		if elapsed < min {
			min = elapsed
		}
		fmt.Println("單一最大花費時間", max)
		fmt.Println("單一最小花費時間", min)

		//sum
		var totalx time.Duration
		var sumNum time.Duration
		for _, valuex := range timeSlice {
			sumNum++
			totalx += valuex
		}
		fmt.Println("壓測總數花費時間", totalx)

		//avg
		fmt.Println("平均花費時間", totalx/sumNum)

		fmt.Println("-----")

	})
	c.Start()
	select {}
}

func apiRoutine(round int) {
	for i := 0; i < round; i++ {
		api()
	}
}

func api() {
	// resp, _ := http.Get("http://weather.json.tw/api")
	resp, _ := http.Get("http://bb.dev.d2/api/domain/32/levels")
	defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// return body
}
