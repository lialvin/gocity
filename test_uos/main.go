package main

import (
	"fmt"
  
    uos "github.com/lialvin/uos-go"
    //"encoding/json"
    "encoding/hex"
	"encoding/json"	
	"os"
	"github.com/lialvin/uos-go/token"
)

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

   
/*
type GetTableRowsRequest struct {
	Code       string `json:"code"` // Contract "code" account where table lives
	Scope      string `json:"scope"`
	Table      string `json:"table"`
	LowerBound string `json:"lower_bound,omitempty"`
	UpperBound string `json:"upper_bound,omitempty"`
	Limit      uint32 `json:"limit,omitempty"`          // defaults to 10 => chain_plugin.hpp:struct get_table_rows_params
	KeyType    string `json:"key_type,omitempty"`       // The key type of --index, primary only supports (i64), all others support (i64, i128, i256, float64, float128, ripemd160, sha256). Special type 'name' indicates an account name.
	Index      string `json:"index_position,omitempty"` // Index number, 1 - primary (first), 2 - secondary index (in order defined by multi_index), 3 - third index, etc. Number or name of index can be specified, e.g. 'secondary' or '2'.
	EncodeType string `json:"encode_type,omitempty"`    // The encoding type of key_type (i64 , i128 , float64, float128) only support decimal encoding e.g. 'dec'" "i256 - supports both 'dec' and 'hex', ripemd160 and sha256 is 'hex' only
	JSON       bool   `json:"json"`                     // JSON output if true, binary if false
}
*/
type UidSystemResp struct {
	uidrows []UidSystem `json:"rows"`
}

// cli  get table uosuidwallet uosuidwallet uidsystem --key-type i64 --index 1   -L testtesttes3 -U testtesttes3  -l 1  
func getAPIURL() string {
	apiURL := os.Getenv("UOS_GO_API_URL")
	if apiURL != "" {
		return apiURL
	}

	return "https://testrpc1.uosio.org:20580"
}

func  Push_transfer_UOS() {
	api := uos.New(getAPIURL())

	keyBag := &uos.KeyBag{}
	err := keyBag.ImportPrivateKey("5JifSQWXDYCdf2qAhfGWpkJM83XmcRuS9uskt5YcUJgyDTrjxhC")
	if err != nil {
		panic(fmt.Errorf("import private key: %s", err))
	}
	api.SetSigner(keyBag)

	from := uos.AccountName("uosuidwallet")
	to := uos.AccountName("ulordusertm2")
	quantity, err := uos.NewUOSAssetFromString("1.0000 UOS")
	memo := "uidback"

	if err != nil {
		panic(fmt.Errorf("invalid quantity: %s", err))
	}

	txOpts := &uos.TxOptions{}
	if err := txOpts.FillFromChain(api); err != nil {
		panic(fmt.Errorf("filling tx opts: %s", err))
	}

	tx := uos.NewTransaction([]*uos.Action{token.NewTransfer(from, to, quantity, memo)}, txOpts)
	signedTx, packedTx, err := api.SignTransaction(tx, txOpts.ChainID, uos.CompressionNone)
	if err != nil {
		panic(fmt.Errorf("sign transaction: %s", err))
	}

	content, err := json.MarshalIndent(signedTx, "", "  ")
	if err != nil {
		panic(fmt.Errorf("json marshalling transaction: %s", err))
	}

	fmt.Println(string(content))
	fmt.Println()

	response, err := api.PushTransaction(packedTx)
	if err != nil {
		panic(fmt.Errorf("push transaction: %s", err))
	}

	fmt.Printf("Transaction [%s] submitted to the network succesfully.\n", hex.EncodeToString(response.Processed.ID))
}


func main() {

    var out []UidSystem

    s1 := uos.GetTableRowsRequest{"uosuidwallet","uosuidwallet" , "uidsystem", "asasas121212", "asasas121212", 1,  "i64", "1", "", true} 
    
    api := uos.New("https://testrpc1.uosio.org:20580")
    //infoResp, _ := api.GetInfo()
    //accountResp, _ := api.GetAccount("uosio")
    inforow1, _err := api.GetTableRows(s1)

    err2:=  inforow1.JSONToStructs( &out )

    Push_transfer_UOS();
    fmt.Println(" uiduoswallet len   ", len(out) )    
    
    
    //fmt.Println("info rest:",string(inforow1.Rows))
    fmt.Println("uiduoswallet table:",out[0],err2, _err )
    //fmt.Println("info rest:",infoResp)
    //fmt.Println("Permission for uosio:", accountResp.Permissions[0].RequiredAuth.Keys)


}
