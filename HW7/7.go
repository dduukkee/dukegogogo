//本週go作業：
//設計一個phone interface
//interface需要設計兩個以上的功能，與兩個以上的func
//並實作兩隻不一樣的phone

package main

import (
	"InterFace7"
)

type iphone1 struct {
	sensor  string
	monitor string
	button  string
}

type iphone2 struct {
	attack string
	run    string
	move   string
}

func (r iphone1) Sensor() string {

	return r.sensor
}

func (r iphone1) Monitor() string {

	return r.monitor
}

func (r iphone1) Button() string {

	return r.button
}

func (r iphone2) Attack() string {

	return r.attack
}

func (r iphone2) Run() string {

	return r.run
}

func (r iphone2) Move() string {

	return r.move
}

func main() {
	iphoneOne := iphone1{sensor: "1的sensor", monitor: "1的monitor", button: "1的button"}
	iphoneTwo := iphone2{attack: "2的sensor", run: "2的monitor", move: "2的button"}
	InterFace7.Unlock(iphoneOne)
	InterFace7.Lock(iphoneOne)
	InterFace7.Fire(iphoneTwo)
	InterFace7.Miss(iphoneTwo)
}
