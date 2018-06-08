package main

import (
	"./otherpackage"
	"fmt"
	"strconv" //Package strconv implements conversions to and from string representations of basic data types.
)

func main() {

	a := int(1)
	b := int32(2)
	c := int64(3)
	d := string("999")
	e := float32(88.8)
	f := float64(99.9)
	x := string("taiwan no.") + otherpackage.IntToStr(a) //ლ(◉◞౪◟◉ )ლ
	
	g := int32(a) + b 
	h := int64(g) + c
	i := f / float64(e)
	o := a + strToInt(d)
	fmt.Printf("g:%d\n", g) 
	fmt.Printf("h:%d\n", h)
	fmt.Printf("i:%f\n", i)
	fmt.Printf("o:%d\n", o)
	fmt.Printf("d:%s\n", x)
}

//字串>整數 Atoi
func strToInt(i string) int {

    res, _ := strconv.Atoi(i)
    return res
}