package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1000)
	// goroutine1
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	// goroutine2
	go func() {
		for {
			a, ok := <-ch

			if !ok {
				fmt.Println("close")

				//原因:channel已被關閉了，但goroutine還在運作持續備給值channel造成噴錯
				//修改:等channel都跑完了，才關閉
				close(ch)

				return
			}

			fmt.Println("a: ", a)
		}

	}()

	fmt.Println("ok")
	time.Sleep(time.Second * 3)
}
