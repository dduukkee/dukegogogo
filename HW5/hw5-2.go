package main

import (
	"fmt"
	"time"
)

func main() {
	//計時
	runtime := time.Now()

	var end int = 0
	for i := 1; i <= 20; i++ {
		result := factory(i)
		end = end + result
		println("目前加總=", end)
	}

	println("結果是=", end)

	//計時結束
	sec := time.Since(runtime)
	fmt.Println("花費: ", sec)

}

func factory(max int) (result int) {
	result = 1
	if max > 0 {
		result = max * factory(max-1)
	}
	return
}
