package main

import (
	"fmt"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	inPath := "keystore/UTC--2016-07-28T14-07-00.162110900Z--13f022d72158410433cbd66f5dd8bf6d2d129924"
	outPath := "outputKey/privateKey"
	password := "password"
	keyjson, e := ioutil.ReadFile(inPath)
	if e != nil {
		panic(e)
	}
	key, e := keystore.DecryptKey(keyjson, password)
	if e != nil {
		panic(e)
	}
	e = crypto.SaveECDSA(outPath, key.PrivateKey)
	if e != nil {
		panic(e)
	}
	fmt.Println("Key saved to:", outPath)
}
