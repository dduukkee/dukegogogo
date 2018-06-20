package main

import (
	"encoding/json"
	"fmt"
	// "reflect"
)

func main() {

	Q1()

	ipList := []string{"192.17.55.3", "192.17.33.17", "192.17.99.52"}
	hostList := []string{"docs.google.com", "ts-phpadmin.vir999.com", "jsonviewer.stack.hu"}
	portList := []int{80, 80, 7711}
	Q2(ipList, hostList, portList)
}

func Q2(ipList []string, hostList []string, portList []int) {
	var containPool []Q2struct

	//2-1
	for i := 0; i < 3; i++ {
		contain := Q2struct{
			IP:   ipList[i],
			HOST: hostList[i],
			PORT: portList[i],
		}

		containPool = append(containPool, contain)
	}
	fmt.Println("2-1 = \n", containPool)

	//2-2 新增 嗯?
	contain := Q2struct{
		IP:   "177.18.2.33",
		HOST: "github.com",
		PORT: 80,
	}
	containPool = append(containPool, contain)
	fmt.Println("2-2 = \n", containPool)

	//2-3 point updata
	updateSetting(&containPool[2].PORT)
	fmt.Println("2-3 = \n", containPool)
}

func updateSetting(point *int) {
	*point = 80

}

type Q2struct struct {
	IP   string
	HOST string
	PORT int
}

func Q1() {
	var result Api
	//過來
	err := json.Unmarshal(getApiResult(), &result)
	if err != nil {
		fmt.Println(err.Error())
	}
	//讓我看看
	fmt.Printf("1-1 id %+v \n", result.Ret[0].ID)
	fmt.Printf("1-1 username %+v \n", result.Ret[0].Username)
	fmt.Printf("1-1 Currency %+v \n", result.Ret[0].Currency)
	fmt.Printf("1-1 Cash %+v \n", result.Ret[0].Cash)
}

func getApiResult() []byte {
	userData := `{"ret":[{"id":148500286,"parent_id":1414663,"parent":1414663,"all_parents":[1414663,59735,58971,24367,6],"domain":6,"last_login":"2018-06-12T09:17:51+0800","currency":"CNY","password_expire_at":"2020-09-04T18:19:02+0800","password_reset":false,"last_bank":54,"username":"wesley01","block":false,"bankrupt":false,"cash":{"id":71814701,"user_id":148500286,"balance":4923642.0545,"pre_sub":0,"pre_add":0,"currency":"CNY","last_entry_at":"20180612155900"},"cash_fake":null,"card":null,"enabled_card":null}],"result":"ok","profile":{"execution_time":14,"server_name":"ipl-web01.rd5.qa"}}`

	return []byte(userData)
}

type Api struct {
	Ret []struct {
		ID               int    `json:"id"`
		ParentID         int    `json:"parent_id"`
		Parent           int    `json:"parent"`
		AllParents       []int  `json:"all_parents"`
		Domain           int    `json:"domain"`
		LastLogin        string `json:"last_login"`
		Currency         string `json:"currency"`
		PasswordExpireAt string `json:"password_expire_at"`
		PasswordReset    bool   `json:"password_reset"`
		LastBank         int    `json:"last_bank"`
		Username         string `json:"username"`
		Block            bool   `json:"block"`
		Bankrupt         bool   `json:"bankrupt"`
		Cash             struct {
			ID          int     `json:"id"`
			UserID      int     `json:"user_id"`
			Balance     float64 `json:"balance"`
			PreSub      int     `json:"pre_sub"`
			PreAdd      int     `json:"pre_add"`
			Currency    string  `json:"currency"`
			LastEntryAt string  `json:"last_entry_at"`
		} `json:"cash"`
		CashFake    interface{} `json:"cash_fake"`
		Card        interface{} `json:"card"`
		EnabledCard interface{} `json:"enabled_card"`
	} `json:"ret"`
	Result  string `json:"result"`
	Profile struct {
		ExecutionTime int    `json:"execution_time"`
		ServerName    string `json:"server_name"`
	} `json:"profile"`
}
