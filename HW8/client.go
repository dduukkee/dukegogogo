package main

import (
	"net/http"
	"fmt"
    "io/ioutil"
)


func main() {
	resp, _ := http.Get("http://localhost:1323/money")
	defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))

    resp2, _ := http.Get("http://localhost:1323/anyththing/moneyyyyyy")
    defer resp2.Body.Close()
    body2, _ := ioutil.ReadAll(resp2.Body)
    fmt.Println(string(body2))

}