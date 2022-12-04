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

	balance, err := instance.BalanceOf(nil, common.HexToAddress("0x9949f7e672a568bB3EBEB777D5e8D1c1107e96E5"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("balance is %d\n", balance)

}
