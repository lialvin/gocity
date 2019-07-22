package main

import (
    
    "fmt"
	"io/ioutil"
	"sort"
    uos "github.com/lialvin/uos-go"
    //"encoding/json"
    "strconv"
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
	Sslice             []int                   `json:"uidsloce"` 
	Startresult        map[string]int          `json:"startresult"` 
}

//根据key排序
 func sortbykey(mp map[string]int) {
	var newMp = make([]string, 0)
	for k, _ := range mp {
	   newMp = append(newMp, k)
	}
	sort.Strings(newMp)
	for _, v := range newMp {
	   fmt.Println("根据key排序后的新集合》》   key:", v, "    value:", mp[v])
	}
 }
 
 //Uid排序结构体
type UidTop struct {
	//姓名
   name  string
   //成绩
   score int
}

type UidTops []UidTop

//Len()
func (s UidTops) Len() int {
	return len(s)
}

//Less():成绩将有低到高排序
func (s UidTops) Less(i, j int) bool {
	return s[i].score > s[j].score
}

//Swap()
func (s UidTops) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}


// 根据value排序
 func sortbyvalue(mp map[string]int) {
	var UidTops UidTops
	temobj :=UidTop{} 
	for oldk, v := range mp {
		temobj.name= oldk 
		temobj.score= v
		UidTops = append(UidTops, temobj)	   
	}
	sort.Sort(UidTops)	
	//sort.Ints(mp.values)
	
	for _, obj:= range UidTops {
	   fmt.Println("根据value排序后的新集合》》  key:",obj.name, " value",obj.score )
	}
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
		return 0
		//return IsAccount(account_name)    
		//panic(fmt.Errorf("get account: %s", err))
	}

	bytes, err := json.Marshal(info)
	if err != nil {
		//return 0xff
		//panic(fmt.Errorf("json marshal response: %s", err))
	}
	_=len(bytes)
	//fmt.Println(string(bytes))
	return 1
}


// 获取账户表 starttime = "156344000"
func  (uidSys *UidSys) GetUosUidAccount(starttime  string ) (int , int ){
    var out []UidAccount

	api := uos.New("https://testrpc1.uosio.org:20580")
    // -L starttime 最小值是 starttime
    sql := uos.GetTableRowsRequest{"uosuidwallet","uosuidwallet" , "uidaccount", starttime, "", 50,  "i64", "2", "", true} 
        
    //infoResp, _ := api.GetInfo()    
    inforow1, _err := api.GetTableRows(sql)
	err2:=  inforow1.JSONToStructs( &out )
	count:= len(out)
	fmt.Println(" uidaccount len   ", count ,_err, err2  )
	maxTime :=0
	repeat:=0
	for  kidx , uidobj := range out {		
		//if IsAccount(uidobj.Uid)== 1 {
		if(true){				 
			_, ok  := uidSys.UidAccountData[uidobj.Uid]		
			if !ok {
				uidSys.UidAccountData[uidobj.Uid] = uidobj	
				uidSys.GetUosUidSystem(uidobj.Uid,uidobj.Creattime );	
						
				if uidobj.Creattime >maxTime{
					maxTime= uidobj.Creattime 
					uidSys.Sslice = append(uidSys.Sslice, maxTime)
				}
			}else{
				repeat++	
			}
		}
		if kidx==4||kidx==0  {
			//fmt.Println(" read uos idx ", kidx, "  ",uidobj.Uid ,"  ", maxTime ," repeat ",repeat )	
		}
	}
		
	return maxTime ,count
}
// 按时间统计 
func (uidSys *UidSys)  StatByTime(startTime int ,endTime int) {
	//UidTimeData[uidobj.Creattime] =  
	idx1:=0       
	keyTime1:=0    
	for  idx, key  := range uidSys.Sslice {
		//fmt.Println(" StatByTime", key, "  ",idx   )	
		_=key
		idx1=idx
		if uidSys.Sslice[idx] >= startTime && uidSys.Sslice[idx] < endTime  {
			keyTime:= uidSys.Sslice[idx] 
			keyTime1 =keyTime 
			for  _, obj := range uidSys.UidTimeData[keyTime] {	
				if(obj.Highuid==""){
					uidSys.Startresult["no"] = uidSys.Startresult["no"]+1 
				}else{
					uidSys.Startresult[obj.Highuid] = uidSys.Startresult[obj.Highuid]+1 
				}				
			}
		}		
	}
	sortbyvalue(uidSys.Startresult)	
	fmt.Println(" Stat time ", idx1 , " time ", keyTime1 ,uidSys.Startresult  )	
}

func (uidSys *UidSys) GetUosUidSystem(account_name string, creattime int  ){
    var out []UidSystem
	account_name=account_name+" "
    s1 := uos.GetTableRowsRequest{"uosuidwallet","uosuidwallet" , "uidsystem",account_name, account_name, 1,  "i64", "1", "", true} 
    
    api := uos.New("https://testrpc1.uosio.org:20580")
    //infoResp, _ := api.GetInfo()
    
    inforow1, _err := api.GetTableRows(s1)

    err2:=  inforow1.JSONToStructs( &out )

	if  len(out) !=1  {
		fmt.Println(" uiduoswallet len err  ", len(out),_err,err2)    	
	}

	uidSys.UidSystemData[account_name] = out[0] 
	uidSys.UidTimeData[creattime] = append(uidSys.UidTimeData[creattime], out[0])
    
        
    //fmt.Println("info rest:",string(inforow1.Rows))
   // fmt.Println("uiduoswallet table:",out[0],err2, _err )
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// map[string][]string  "D:\\uidaccountdata.txt"
func (uidSys *UidSys)  ReadUidAccount(filename string) int {
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
	fmt.Print(" ReadUidAccount   : ", " record number: ", len(uidSys.UidAccountData)," \n" )
	fmt.Println("max key is :",maxkey)
	// 从 maxkey +1 查找新的key 
	return  maxkey
}

// map[string][]string  "D:\\uidaccountdata.txt"
func (uidSys *UidSys)  ReadUidSystem(filename string) int {
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
	err = json.Unmarshal([]byte(string(buf)), &uidSys.UidSystemData)	
    if err != nil {
        panic(err)
	}

	fmt.Print(" ReadUidSystem   : ", " record number: ", len(uidSys.UidSystemData)," \n" )
	return 0
}

// map[string][]string  "D:\\uidaccountdata.txt"
func (uidSys *UidSys)  ReadObjTime(filename string) int {
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
	err = json.Unmarshal([]byte(string(buf)), &uidSys.UidTimeData)	
    if err != nil {
        panic(err)
	}
	fmt.Print(" ReadObjTime   : ", " record number: ", len(uidSys.UidTimeData)," \n" )
	return 0
}
// map[string][]string  "D:\\uidaccountdata.txt"
func (uidSys *UidSys)  ReadSslice(filename string) int {
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
	err = json.Unmarshal([]byte(string(buf)), &uidSys.Sslice)	
    if err != nil {
        panic(err)
	}
	fmt.Print(" ReadSslice   : ", " record number: ", len(uidSys.Sslice)," \n" )
	return 0
}

func (uidSys *UidSys) WriteUidAccount( filename string) {
	//var buf []byte 
	fmt.Print(" WriteUidAccount   : ",filename ," record number: ", len(uidSys.UidAccountData)," \n" )
	buf , _err := json.Marshal( uidSys.UidAccountData )
	
	if  _err != nil  {
		fmt.Print(" Marshal UidAccountData err: ", _err)
	}
	     
    err := ioutil.WriteFile(filename,buf,os.ModeAppend)
    if err != nil {
        fmt.Print(err)
	}	
}

func (uidSys *UidSys) WriteUidSystem( filename string) {
	//var buf []byte 
	fmt.Print(" WriteUidSystem : ",filename ," record number: ", len(uidSys.UidSystemData)," \n" )
	buf , _err := json.Marshal( uidSys.UidSystemData )
	
	if  _err != nil  {
		fmt.Print(" Marshal UidAccountData err: ", _err)
	}
	     
    err := ioutil.WriteFile(filename,buf,os.ModeAppend)
    if err != nil {
        fmt.Print(err)
	}	
}

func (uidSys *UidSys) WriteObjTime( filename string) {
	//var buf []byte 
	fmt.Print(" WriteObjTime   : ",filename ," record number: ", len(uidSys.UidTimeData)," \n" )
	buf , _err := json.Marshal( uidSys.UidTimeData )
	
	if  _err != nil  {
		fmt.Print(" Marshal UidAccountData err: ", _err)
	}
	     
    err := ioutil.WriteFile(filename,buf,os.ModeAppend)
    if err != nil {
        fmt.Print(err)
	}	
}

func (uidSys *UidSys) WriteSslice( filename string) {
	//var buf []byte 
	fmt.Print(" WriteSslice   : ",filename ," record number: ", len(uidSys.Sslice)," \n" )
	buf , _err := json.Marshal( uidSys.Sslice )
	
	if  _err != nil  {
		fmt.Print(" Marshal UidAccountData err: ", _err)
	}
	     
    err := ioutil.WriteFile(filename,buf,os.ModeAppend)
    if err != nil {
        fmt.Print(err)
	}	
}

func (uidSys *UidSys) ReadfromFile(  ) {
	uidSys.ReadUidAccount("D:\\uidsystemdata.txt")   
	uidSys.ReadUidSystem("D:\\uidsystemdatb.txt") 
	uidSys.ReadObjTime("D:\\uidsystemdatc.txt") 
	uidSys.ReadSslice("D:\\uidsystemdatd.txt") 
}

func (uidSys *UidSys) WritetoFile(  ) {
	uidSys.WriteUidAccount("D:\\uidsystemdata.txt")   
	uidSys.WriteUidSystem("D:\\uidsystemdatb.txt") 
	uidSys.WriteObjTime("D:\\uidsystemdatc.txt") 
	uidSys.WriteSslice("D:\\uidsystemdatd.txt") 
}
// 从网络读取
func (uidSys *UidSys)  SyncUid(startTime  string  ){
	//var buf []byte 
	newUidTime,count := uidSys.GetUosUidAccount("1000")
	for {   
		if count <5{
			break;	
		} 
		newstarttime := strconv.Itoa(newUidTime)   //数字变成字符串
		newUidTime,count  = uidSys.GetUosUidAccount(newstarttime)	
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
