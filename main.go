package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/dlclark/regexp2"
	addr "github.com/fbsobreira/gotron-sdk/pkg/address"
)

func GenerateKey() (wif string, address string) {
	priv, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return "", ""
	}
	if len(priv.D.Bytes()) != 32 {
		for {
			priv, err := btcec.NewPrivateKey(btcec.S256())
			if err != nil {
				continue
			}
			if len(priv.D.Bytes()) == 32 {
				break
			}
		}
	}
	a := addr.PubkeyToAddress(priv.ToECDSA().PublicKey)
	address = a.String()
	wif = hex.EncodeToString(priv.D.Bytes())
	return
}

func judge(str string) bool {
	var validID = regexp2.MustCompile(`^T.*([\w])\1{3,}$`, 0)
	res, err := validID.MatchString(str)
	if err != nil {
		return false
	}
	return res
}

func gen(str string) {
	for true {
		seed, address := GenerateKey()
		if judge(address) {
			fmt.Println("seed:", seed)
			fmt.Println("address:", address)
			fmt.Println("-----------------------------------------------------")
		}
	}
}

func main() {
	var pool int
	var str string
	flag.IntVar(&pool, "t", 4, "Thread Count")
	flag.Parse()
	fmt.Println("-----------------------------------------------------")
	for i := 0; i < pool; i++ {
		go gen(str)
	}
	for {

	}
}
