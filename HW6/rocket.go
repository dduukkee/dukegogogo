package main

import (
	"fmt"
	"time"
)

type rocket struct {
	time time.Duration
}

type RocketInterface interface {
	gogo() bool
}

////////////
func (r rocket) gogo() bool {
	for i := r.time; i > 0; i-- {
		fmt.Printf("倒數%v秒發射\n", i)
		time.Sleep(time.Duration(1) * time.Second)
	}
	return true
}

func main() {
	rocket := rocket{time: 3}
	RunRocket(rocket)
}

func RunRocket(c RocketInterface) {
	fmt.Printf("發射測試 %v \n", c.gogo())
}
