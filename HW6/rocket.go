package main

import (
	"RInt"
	"fmt"
	"time"
)

type rocket1 struct {
	time time.Duration
}

type rocket2 struct {
	time time.Duration
}

func (r rocket1) Gogo() bool {
	for i := r.time; i > 0; i-- {
		fmt.Printf("一號倒數%v秒發射\n", i)
		time.Sleep(time.Duration(1) * time.Second)
	}
	return true
}

func (r rocket2) Gogo() bool {
	for i := r.time; i > 0; i-- {
		fmt.Printf("二號倒數%v秒發射\n", i)
		time.Sleep(time.Duration(1) * time.Second)
	}
	return true
}

func main() {
    rocket1 := rocket1{time: 3}
    rocket2 := rocket2{time: 5}
    RInt.Launch(rocket1)
	RInt.Launch(rocket2)

}
