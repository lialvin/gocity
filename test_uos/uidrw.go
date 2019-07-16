package main

import (
    
    "fmt"
	"io/ioutil"
	  
    uos "github.com/lialvin/uos-go"
    //"encoding/json"
    
	"encoding/json"	
	"os"
)

type UidAccount struct {
	Uid           string         `json:"uid"`
	Account       string         `json:"account"`
	Creattime     int            `json:"creattime"`
	Spentmoney    int            `json:"spentmoney"`
	Storage       int            `json:"storage"`
	Ispayed       bool           `json:"ispayed"`
}


type UidSystem struct {
	Uid           string         `json:"uid"`
	Time          int            `json:"time"`
	Privilege     int            `json:"privilege"`
	Uidflag       int            `json:"uidflag"`
	Childuidnum   int            `json:"childuidnum"`
	Recommenduid  string         `json:"recommenduid"`
	Prevuid       string         `json:"prevuid"`
    Nextuid       string         `json:"nextuid"`
    Lastchilduid  string         `json:"lastchilduid"`
    Highuid       string         `json:"highuid"`
}


type  UidSys   struct {    
	UidSystemData      map[string]UidSystem   `json:"uidsystemdata"`
	UidAccountData     map[int][]UidAccount     `json:"uidaccountdata"` 
}


func IsAccount( ) {
	api := uos.New(getAPIURL())

	account := uos.AccountName("uosuidwallet")
	info, err := api.GetAccount(account)
	if err != nil {
		if err == uos.ErrNotFound {
			fmt.Printf("unknown account: %s", account)
			return
		}

		panic(fmt.Errorf("get account: %s", err))
	}

	bytes, err := json.Marshal(info)
	if err != nil {
		panic(fmt.Errorf("json marshal response: %s", err))
	}

	fmt.Println(string(bytes))
}

// 获取账户表
func  (uidSys *UidSys) GetUosUidAccount(starttime  string ){

    var out []UidAccount

	api := uos.New("https://testrpc1.uosio.org:20580")
    // -L starttime 最小值是 starttime
    sql := uos.GetTableRowsRequest{"uosuidwallet","uosuidwallet" , "uidaccount", starttime, "", 20,  "i64", "2", "", true} 
        
    //infoResp, _ := api.GetInfo()    
    inforow1, _err := api.GetTableRows(sql)
	err2:=  inforow1.JSONToStructs( &out )
	
	fmt.Println(" uidaccount len   ", len(out) ,_err, err2  )
	//for( )
	
}

func (uidSys *UidSys) GetUosUidSystem(){
    var out []UidSystem

    s1 := uos.GetTableRowsRequest{"uosuidwallet","uosuidwallet" , "uidsystem", "asasas121212", "asasas121212", 1,  "i64", "1", "", true} 
    
    api := uos.New("https://testrpc1.uosio.org:20580")
    //infoResp, _ := api.GetInfo()
    IsAccount();
    inforow1, _err := api.GetTableRows(s1)

    err2:=  inforow1.JSONToStructs( &out )


    fmt.Println(" uiduoswallet len   ", len(out) )    
        
    //fmt.Println("info rest:",string(inforow1.Rows))
    fmt.Println("uiduoswallet table:",out[0],err2, _err )
}

// map[string][]string
func (uidSys *UidSys)  ReadfromFile() {
    b, err := ioutil.ReadFile("D:\\whitelist.csv")
    if err != nil {
		fmt.Print(err)
		return 
    }

    whitelist := map[string]bool{}
    err = json.Unmarshal([]byte(string(b)), &whitelist)
    if err != nil {
        panic(err)
    }

    for key, value := range whitelist {
        fmt.Println("key", key, "value", value)
        fmt.Printf("%T", value)
    }
}

func (uidSys *UidSys) WritetoFile( data [] byte) {
    err := ioutil.WriteFile("D:\\whitelist.csv",data,os.ModeAppend)
    if err != nil {
        fmt.Print(err)
    }

    whitelist := map[string]bool{}
    //err = json.Unmarshal([]byte(string(b)), &whitelist)
    if err != nil {
        panic(err)
    }

    for key, value := range whitelist {
        fmt.Println("key", key, "value", value)
        fmt.Printf("%T", value)
    }
}


