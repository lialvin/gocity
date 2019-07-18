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
	UidSystemData      map[string]UidSystem    `json:"uidsystemdata"`
	UidAccountData     map[string]UidAccount   `json:"uidaccountdata"` 
	UidTimeData        map[int][]UidSystem     `json:"uidtimedata"` 
	sslice             []int                   `json:"uidsloce"` 
	startresult        map[string]int          `json:"startresult"` 
}


func IsAccount( account_name string )  int{
	api := uos.New(getAPIURL())

	account := uos.AccountName(account_name)
	info, err := api.GetAccount(account)
	if err != nil {
		if err == uos.ErrNotFound {
			fmt.Printf("unknown account: %s", account_name)
			return 0
		}

		return IsAccount(account_name)    
		//panic(fmt.Errorf("get account: %s", err))
	}

	bytes, err := json.Marshal(info)
	if err != nil {
		//return 0xff
		//panic(fmt.Errorf("json marshal response: %s", err))
	}
	
	fmt.Println(string(bytes))
	return 1
}


// 获取账户表 starttime = "156344000"
func  (uidSys *UidSys) GetUosUidAccount(starttime  string ) (int , int ){
    var out []UidAccount

	api := uos.New("https://testrpc1.uosio.org:20580")
    // -L starttime 最小值是 starttime
    sql := uos.GetTableRowsRequest{"uosuidwallet","uosuidwallet" , "uidaccount", starttime, starttime, 50,  "i64", "2", "", true} 
        
    //infoResp, _ := api.GetInfo()    
    inforow1, _err := api.GetTableRows(sql)
	err2:=  inforow1.JSONToStructs( &out )
	count:= len(out)
	fmt.Println(" uidaccount len   ", count ,_err, err2  )
	maxTime :=0
	for  _ , uidobj := range out {
		if IsAccount(uidobj.Uid)== 1 {			
			_, ok  := uidSys.UidAccountData[uidobj.Uid]		
			if !ok {
				uidSys.UidAccountData[uidobj.Uid] = uidobj	
				uidSys.GetUosUidSystem(uidobj.Uid,uidobj.Creattime );					
				if uidobj.Creattime >maxTime{
					maxTime= uidobj.Creattime 
					uidSys.sslice = append(uidSys.sslice, maxTime)
				}
			}	
		}
	}
	return maxTime ,count
}
// 按时间统计 
func (uidSys *UidSys)  StatByTime(startTime int ,endTime int) {
	//UidTimeData[uidobj.Creattime] =  
	for   _, idx := range uidSys.sslice {
		if uidSys.sslice[idx] >= startTime && uidSys.sslice[idx] < endTime  {
			keyTime:= uidSys.sslice[idx] 
			for  _,obj := range uidSys.UidTimeData[keyTime] {				
				uidSys.startresult[obj.Highuid] = uidSys.startresult[obj.Highuid]+1 
			}
		}		
	}
}

func (uidSys *UidSys) GetUosUidSystem(account_name string, creattime int  ){
    var out []UidSystem

    s1 := uos.GetTableRowsRequest{"uosuidwallet","uosuidwallet" , "uidsystem",account_name, account_name, 1,  "i64", "1", "", true} 
    
    api := uos.New("https://testrpc1.uosio.org:20580")
    //infoResp, _ := api.GetInfo()
    
    inforow1, _err := api.GetTableRows(s1)

    err2:=  inforow1.JSONToStructs( &out )

	if  len(out) !=1  {
		fmt.Println(" uiduoswallet len err  ", len(out) )    	
	}

	uidSys.UidSystemData[account_name] = out[0] 
	uidSys.UidTimeData[creattime] = append(uidSys.UidTimeData[creattime], out[0])
    
        
    //fmt.Println("info rest:",string(inforow1.Rows))
    fmt.Println("uiduoswallet table:",out[0],err2, _err )
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// map[string][]string  "D:\\uidaccountdata.txt"
func (uidSys *UidSys)  ReadfromFile(filename string) int {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
		fmt.Println("file exist start read :",exist)
		return  0
	}
	
    buf, err := ioutil.ReadFile(filename)
    if err != nil {
		fmt.Print(err)
		return  0
    }

    //whitelist := map[string]bool{}
	//err = json.Unmarshal([]byte(string(buf)), &whitelist)
	err = json.Unmarshal([]byte(string(buf)), &uidSys.UidAccountData)	
    if err != nil {
        panic(err)
	}
	
	var maxkey =0 
	
	for  _  , key    := range uidSys.UidAccountData {
		if key.Creattime >maxkey {
		    maxkey= key.Creattime
		}		
	}
	fmt.Println("max key is :",maxkey)
	// 从 maxkey +1 查找新的key 
	return  maxkey
}

func (uidSys *UidSys) WritetoFile( filename string) {
	//var buf []byte 
	buf , _err := json.Marshal( uidSys.UidAccountData )
	
	if  _err != nil  {
		fmt.Print(" Marshal UidAccountData err: ", _err)
	}
	     
    err := ioutil.WriteFile(filename,buf,os.ModeAppend)
    if err != nil {
        fmt.Print(err)
    }
}
// 从网络读取
func (uidSys *UidSys)  SyncUid(startTime  string  ){
	//var buf []byte 
	newUidTime,count := uidSys.GetUosUidAccount("1000")
	for {   
		if count <50{
			break;	
		} 
		newUidTime,count  = uidSys.GetUosUidAccount(string(newUidTime))	
	}
		

	fmt.Println("max time is :",newUidTime)
}


func (uidSys *UidSys)  printuid(uidstd string){

	var out []UidSystem
	
	uid:=uidstd+" "
    s1 := uos.GetTableRowsRequest{"uosuidwallet","uosuidwallet" , "uidsystem",uid, uid, 1,  "i64", "1", "", true} 
    //s1 := uos.GetTableRowsRequest{"uosuidwallet","uosuidwallet" , "uidsystem","", "", 10,  "i64", "1", "", true} 
    api := uos.New("https://testrpc1.uosio.org:20580")
    //infoResp, _ := api.GetInfo()
    
    inforow1, _err := api.GetTableRows(s1)
	fmt.Println("uidsystem test uid:",string(inforow1.Rows) )
    err2:=  inforow1.JSONToStructs( &out )
	
	if  len(out) !=1  {
		fmt.Println(" uidsystem   len=0,  find err  ",uid, len(out),_err ,err2)    	
		return 
	}   
        
 
	fmt.Println("\n\n")
}

func (uidSys *UidSys)  StatUid(){

	
}
