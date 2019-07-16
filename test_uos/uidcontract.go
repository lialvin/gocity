

package main

import (
	"fmt"
  
    uos "github.com/lialvin/uos-go"
    //"encoding/json"
    "encoding/hex"
	"encoding/json"		
	"github.com/lialvin/uos-go/token"
)


// cli  get table uosuidwallet uosuidwallet uidsystem --key-type i64 --index 1   -L testtesttes3 -U testtesttes3  -l 1  
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

	fmt.Println(len(string(content)))
	fmt.Println()

	response, err := api.PushTransaction(packedTx)
	if err != nil {
		panic(fmt.Errorf("push transaction: %s", err))
	}

	fmt.Printf("Transaction [%s] submitted to the network succesfully.\n", hex.EncodeToString(response.Processed.ID))
}