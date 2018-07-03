package main

import (
	"fmt"
	"sync"
	"time"
)

//伐廠每秒產1頓紙漿
//紙廠每秒可個別處 0.5頓與0.3頓的紙漿，製成5000張與3000張紙
//有間印刷廠每秒可印6000張紙
//請把這三種廠成產鏈，讓印刷廠總共印製60000張紙

type wooderrrrr struct {
	produceWooderKg float64 //每秒產X頓紙漿
	color           int     //訊息顯示的顏色
}

type paperrrr struct {
	useWooderKg  float64 //需要幾頓
	producePaper int     //每秒產X頓紙漿
	color        int     //訊息顯示的顏色
}

type printerrrrr struct {
	printNum int //製成張數
	color    int //訊息顯示的顏色
}

func main() {
	//伐廠 每秒產1頓紙漿
	wooderPower := wooderrrrr{
		produceWooderKg: 1, //每秒產1頓紙漿
		color:           33,
	}

	//紙廠1 每秒處理 0.5頓，製成5000張
	paperPower1 := paperrrr{
		useWooderKg:  0.5,
		producePaper: 5000,
		color:        20,
	}
	//紙廠2 每秒處理 0.3頓，製成3000張
	paperPower2 := paperrrr{
		useWooderKg:  0.3,
		producePaper: 3000,
		color:        20,
	}

	//印刷廠 每秒可印6000張紙
	printPower := printerrrrr{
		printNum: 6000,
		color:    34,
	}

	//建立WaitGroup
	var wg sync.WaitGroup
	//建立互斥鎖
	// var mutex sync.Mutex

	allWood := float64(0) //目前有多少木頭
	var allPaper = 0      //目前有多少紙
	var allPrint = 0      //目前印了多少紙
	//var maxPrint = 60000 //總共要印製60000張紙

	wg.Add(2)
	go wooderFunc(wooderPower, &allWood, &allPrint)
	go paperFunc(paperPower1, &allWood, &allPaper, &allPrint, &wg, 1)
	go paperFunc(paperPower2, &allWood, &allPaper, &allPrint, &wg, 2)
	go printerFunc(printPower, &allPaper, &allPrint)

	//等待所有工作完成
	wg.Wait()

}

func wooderFunc(w wooderrrrr, allWood *float64, allPrint *int) {

	//印好60000紙之前都給我好好去伐木木木木木
	for i := *allPrint; i < 60000; i = *allPrint {
		//工作前的數量
		beforeWoodNum := *allWood

		//木總產量++
		*allWood = *allWood + w.produceWooderKg

		fmt.Printf("%c[0;40;%dm伐木場:目前有%v噸木頭，累計生產了%v噸木頭\n%c[0m", 0x1B, w.color, beforeWoodNum, *allWood, 0x1B)

		//休息
		time_s := time.Duration(1) * time.Second
		time.Sleep(time_s)
	}
}

func paperFunc(p paperrrr, allWood *float64, allPaper *int, allPrint *int, wg *sync.WaitGroup, who int) {

	//印好60000紙之前都給我好好去造紙紙紙紙紙紙紙
	for i := *allPrint; i < 60000; i = *allPrint {
		//工作前的木頭數量
		beforeWoodNum := *allWood

		//若木頭有超過可拿取的量 就去生產紙
		if *allWood >= p.useWooderKg {
			*allPaper = *allPaper + p.producePaper //紙+
			*allWood = *allWood - p.useWooderKg    //木頭-
			fmt.Printf("%c[0;40;%dm紙廠%v:目前有%v噸木頭用，累計產了%v張紙，剩下%v噸木頭\n%c[0m", 0x1B, p.color, who, beforeWoodNum, *allPaper, *allWood, 0x1B)
		} else {
			//阿 木頭不夠
			fmt.Printf("%c[0;40;%dm紙廠%v:目前剩%v噸木頭可用，不夠，休息\n%c[0m", 0x1B, p.color, who, beforeWoodNum, 0x1B)
		}

		time_s := time.Duration(1) * time.Second
		time.Sleep(time_s)
	}

	//工作完成 回報WaitGroup -1
	wg.Done()
}

func printerFunc(p printerrrrr, allPaper *int, allPrint *int) {

	//印好60000紙之前都給我好好去印啦印啦印啦印啦印啦印啦
	for i := *allPrint; i < 60000; i = *allPrint {
		//工作前的 紙張數量
		beforePaperNum := *allPaper
		//木頭有超過可拿取量 就去印印印印印紙
		if *allPaper > p.printNum {
			*allPrint = *allPrint + p.printNum //印刷紙+
			*allPaper = *allPaper - p.printNum //紙-
			fmt.Printf("%c[0;40;%dm印刷廠:目前有%v張紙可用，累計印了%v張紙，剩下%v張紙\n%c[0m", 0x1B, p.color, beforePaperNum, *allPrint, *allPaper, 0x1B)
		} else {
			//阿 紙不夠
			fmt.Printf("%c[0;40;%dm印刷廠:目前只剩有%v張紙，紙不夠，休息\n%c[0m", 0x1B, p.color, beforePaperNum, 0x1B)
		}

		time_s := time.Duration(1) * time.Second
		time.Sleep(time_s)
	}

}
