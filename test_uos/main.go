package main

import (
	"fmt"
	"os"

)

func getAPIURL() string {
	apiURL := os.Getenv("UOS_GO_API_URL")
	if apiURL != "" {
		return apiURL
	}

	return "https://testrpc1.uosio.org:20580"
}


func main() {

    fmt.Println("start read from file ")
    //fmt.Println("info rest:",infoResp)
    //fmt.Println("Permission for uosio:", accountResp.Permissions[0].RequiredAuth.Keys)
    
    uidSysObj := new(UidSys)
    uidSysObj.ReadfromFile()
    // map[string][]UidSystem , map[int][]UidAccount

}

// 增加新功能，统计 新增用户 。 一期(区块周)  每一期 为 7*24*3600  . 统计团体新增用户 从某一秒开始统计   
// 读取表， 生成记录 本地存储  map[] [] 
