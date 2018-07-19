package main

import (
	// "fmt"
	"time"
)

// type CellPhone interface {
// 	Name() string
// 	Size() int
// 	TalkTime() time.Duration
// }

// func (i Iphone) Name() string {
// 	return i.version
// }

// func (i Iphone) Size() int {
// 	return i.height * i.width
// }

// func (i Iphone) TalkTime() time.Duration {
// 	return i.battery * time.Hour
// }

// type Iphone struct {
// 	version       string
// 	width, height int
// 	battery       time.Duration
// }

// //他是圓形的
// type IWatch struct {
// 	version string
// 	radius  int
// 	battery time.Duration
// }

// func main() {
// 	iphone := Iphone{width: 30, height: 90, battery: 24, version: "iphone-XX"}
// 	//pixel := Iphone{width:40,height:120,battery:1300,version:"pixel-3"}

// 	ShowPhone(iphone)
// 	//ShowPhone(pixel)
// }

// func ShowPhone(c CellPhone) {
// 	fmt.Printf("Product %v \n", c.Name())
// 	fmt.Printf("Size %v \n", c.Size())
// 	fmt.Printf("Talk time %v \n", c.TalkTime())
// 	fmt.Println()
// }

func (i iwatch) Name() string {
	return i.version
}

func (i iwatch) Size() int {
	return i.radius * i.radius * 3
}

func (i iwatch) TalkTime() time.Duration {
	return i.battery * time.Hour
}

type popoface interface {
	name() string
	size() int
	TalkTime() time.Duration
}

func (i popoface) factory() {
	a := i.name()
	i.size()
	i.TalkTime()
}

//他是圓形的
type iwatch struct {
	version string
	radius  int
	battery time.Duration
}

type iohweiya struct {
	version       string
	width, height int
	battery       time.Duration
}

func main() {
	// iphone := Iphone{width: 30, height: 90, battery: 24, version: "iphone-XX"}
	//pixel := Iphone{width:40,height:120,battery:1300,version:"pixel-3"}
	iwatch := iwatch{version: 30, radius: 90, battery: 24, version: "iwatch-XX"}

	factory(iwatch)
	//factory(iohweiya)
}
