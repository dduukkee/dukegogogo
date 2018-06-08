package main

import (
	"fmt"
)

//先建立一個長度為 5 的陣列來儲存這些測驗成績，接著將分數填入每個元素中。
//再來使用一個迴圈來計算成績的總和。
//最後我們將成績的總和除以元素的數量，以取得平均值。

func main() {

	var score [5]int = [5]int{11, 22, 33, 44, 55}
	sum := int(0)
	for _, value := range score {
		//fmt.Println("score[", key, "] =", value)
		sum += value
	}
	avg := int(sum / len(score))
	fmt.Println("總和", sum)
	fmt.Println("平均", avg)

	//寫個程式找出串列中最小的數字：
	xarray := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	min := xarray[0]
	for _, xvalue := range xarray {
		if min > xvalue {
			min = xvalue
		}
	}
	fmt.Println("最小:", min)
}
