package main

import (
	"fmt"
	"time"
)

type Iphone struct {
	version       string
	width, height int
	battery       time.Duration
}

//他是圓形的
type IWatch struct {
	version string
	radius  int
	battery time.Duration
}

type CellPhone interface {
	Name() string
	Size() int
	TalkTime() time.Duration
}

////////////
func (i Iphone) Name() string {
	return i.version
}

func (i Iphone) Size() int {
	return i.height * i.width
}

func (i Iphone) TalkTime() time.Duration {
	return i.battery * time.Hour
}

////////////
func (i IWatch) Name() string {
	return i.version
}

func (i IWatch) Size() int {
	return i.radius * i.radius * 3
}

func (i IWatch) TalkTime() time.Duration {
	return i.battery * time.Hour
}

func main() {
	iphone := Iphone{width: 100, height: 200, battery: 24, version: "iphone-XX"}

	IWatch := IWatch{radius: 3, battery: 24, version: "IWatch-dd"}
	ShowIWatch(IWatch)
	ShowPhone(iphone)
}

func ShowIWatch(c IWatch) {
	fmt.Printf("Product %v \n", c.Name())
	fmt.Printf("Size %v \n", c.Size())
	fmt.Printf("Talk time %v \n", c.TalkTime())
	fmt.Println()
}

func ShowPhone(c CellPhone) {
	fmt.Printf("Product %v \n", c.Name())
	fmt.Printf("Size %v \n", c.Size())
	fmt.Printf("Talk time %v \n", c.TalkTime())
	fmt.Println()
}
