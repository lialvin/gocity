package main

import (
	"fmt"

    uos "github.com/lialvin/uos-go"
)


func main() {

    api := uos.New("http://testnet1.eos.io")
    infoResp, _ := api.GetInfo()
    accountResp, _ := api.GetAccount("initn")
    fmt.Println("info rest:",infoResp)
    fmt.Println("Permission for initn:", accountResp.Permissions[0].RequiredAuth.Keys)
 
}
