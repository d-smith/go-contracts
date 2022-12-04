package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/d-smith/go-contract/token"
)

func main() {

	//Client connection
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	// Private key of caller
	privateKey, err := crypto.HexToECDSA("cb1a18dff8cfcee16202bf86f1f89f8b3881107b8192cd06836fda9dbc0fde1b")
	if err != nil {
		log.Fatal(err)
	}

	// Build transactor
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// Get an instance of the smart contract
	address := common.HexToAddress("0xdb98a5bfba239000213813b2615b8a96e950a79b")
	instance, err := token.NewToken(address, client)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := instance.Transfer(auth,
		common.HexToAddress("0x9949f7e672a568bB3EBEB777D5e8D1c1107e96E5"),
		big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
}
