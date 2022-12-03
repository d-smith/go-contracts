package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/d-smith/go-contract/token"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0xdb98a5bfba239000213813b2615b8a96e950a79b")
	instance, err := token.NewToken(address, client)
	if err != nil {
		log.Fatal(err)
	}

	totalSupply, err := instance.TotalSupply(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(totalSupply) // "1.0"
}
