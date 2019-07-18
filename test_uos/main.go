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



func main() {    
    
    // fmt.Printf("Your name is %s", input)
    uidSysObj := new(UidSys)
    for {
        display()
        inputReader := bufio.NewReader(os.Stdin)
        input, _ := inputReader.ReadString('\n')    
        keyval:= strings.TrimSpace(input)
        //fmt.Println("input", input )
        switch keyval {
            case "1" :
                fmt.Println("read from file !")
                uidSysObj.ReadfromFile("D:\\uidsystemdata.txt")                   
            case "2":
                fmt.Println("write to  file  !")
                uidSysObj.WritetoFile("D:\\uidsystemdata.txt")                                
            case "3":
                fmt.Println("sync from uos node")
                uidSysObj.SyncUid("1000")
            case "4":
                startTime:= getTime()
                endTime:= getTime()
                fmt.Println("stat from uos node")                
                uidSysObj.StatByTime(startTime ,endTime )               
            case "5":
                uid:=getinput()
                uidSysObj.printuid(uid)
            case "6":    
                fmt.Println("exit sysyem")
                uidSysObj.WritetoFile("D:\\uidsystemdata.txt") 
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
