package main

import (
	"fmt"

    uos "github.com/lialvin/uos-go"
)


func main() {

    api := uos.New("https://testrpc1.uosio.org:20580")
    infoResp, _ := api.GetInfo()
    accountResp, _ := api.GetAccount("uosio")
    fmt.Println("info rest:",infoResp)
    fmt.Println("Permission for initn:", accountResp.Permissions[0].RequiredAuth.Keys)
    
 
}
