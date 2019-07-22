package main

import (
	"fmt"
    "os"
    "bufio"
    "strconv"
    "strings" 
)

func getAPIURL() string {
	apiURL := os.Getenv("UOS_GO_API_URL")
	if apiURL != "" {
		return apiURL
	}

	return "https://testrpc1.uosio.org:20580"
}


func  display() {
    fmt.Printf("read from file: 1 \n")
    fmt.Printf("write to  file : 2 \n")
    fmt.Printf("sync from uos node: 3 \n")   
    fmt.Printf("statistic uid top 10: 4 \n")    
    fmt.Printf("print uid info: 5 \n")    
    fmt.Printf("exit : 6 \n")    
}
 
func getinput() string{
    inputReader := bufio.NewReader(os.Stdin)
    input, _ := inputReader.ReadString('\n')    
    keyval:= strings.TrimSpace(input)
    return keyval
}


func  getTime()  int {
    keyval :=getinput()
    time, _:=  strconv.Atoi(keyval) 
    return time 
}

/*UidSystemData      map[string]UidSystem    `json:"uidsystemdata"`
UidAccountData     map[string]UidAccount   `json:"uidaccountdata"` 
UidTimeData        map[int][]UidSystem     `json:"uidtimedata"` 
sslice             []int                   `json:"uidsloce"` 
startresult        map[string]int          `json:"startresult"` */

type  testa   struct{

    UidAccountData     map[string]UidAccount   `json:"uidaccountdata"`
   // UidAccount     map[string]UidAccount   `json:"uidaccountdata"`
}


func main() {    
        
    // fmt.Printf("Your name is %s", input)
    //var sss UidAccount

    //uidtest := testa{  ( map[string]UidAccount ) }
    //uidtest.UidAccountData["23"]= sss

    uidSysObj := UidSys{ UidSystemData: map[string]UidSystem{} ,UidAccountData:map[string]UidAccount{} ,UidTimeData: map[int][]UidSystem {}, Startresult: map[string]int{} }
    // UidSystemData =make( map[string]UidSystem ) ,
    //              UidAccountData=make( map[string]UidSystem ),UidTimeData= make( map[int][]UidSystem),startresult= make(map[string]int )   }
    //uidSysObj.SyncUid("1000")
    //uidSysObj.ReadfromFile()   
    //uidSysObj.StatByTime( 1000, 1793267809 )  
    for {
        display()
        inputReader := bufio.NewReader(os.Stdin)
        input, _ := inputReader.ReadString('\n')    
        keyval:= strings.TrimSpace(input)
        //fmt.Println("input", input )
        switch keyval {
            case "1" :
                fmt.Println("read from file !")
                uidSysObj.ReadfromFile()   
            case "2":
                fmt.Println("write to  file  !")
                uidSysObj.WritetoFile( )                                
            case "3":
                fmt.Println("sync from uos node")
                uidSysObj.SyncUid("1000")
            case "4":
                startTime:= getTime()
                endTime:= getTime()
                fmt.Println("stat from uos node",startTime, endTime )        
                //uidSysObj.StatByTime( 1503176126, 1693267809 )        
                uidSysObj.StatByTime(startTime ,endTime )               
            case "5":
                uid:=getinput()
                uidSysObj.printuid(uid)
            case "6":    
                fmt.Println("exit sysyem")
                //uidSysObj.WritetoFile( ) 
                os.Exit(0)
        }
    }


    fmt.Println("start read from file ")
    //fmt.Println("info rest:",infoResp)
    //fmt.Println("Permission for uosio:", accountResp.Permissions[0].RequiredAuth.Keys)  
    
    // map[string][]UidSystem , map[int][]UidAccount

    // 增加新功能，统计 新增用户 。 一期(区块周)  每一期 为 7*24*3600  . 统计团体新增用户 从某一秒开始统计   
    // 读取表， 生成记录 本地存储  map[] [] 
}
