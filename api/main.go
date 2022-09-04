package main

import (
	"JY8752/kubernetes-privatenet-hello-contract/env"
	"JY8752/kubernetes-privatenet-hello-contract/hello"

	// "context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// ctx := context.Background()

	//client
	cl, err := ethclient.Dial(env.RpcUrl())
	if err != nil {
		log.Fatalf("ethclientの初期化に失敗しました. err: %v", err)
	}

	//hello contract
	address := common.HexToAddress(env.ContractAddress())
	hello, err := hello.NewHello(address, cl)
	if err != nil {
		log.Fatalf("hello contractの取得に失敗しました. err: %v", err)
	}

	res, err := hello.Hello(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("helloメソッドの呼び出しに失敗しました. err: %v", err)
	}

	fmt.Printf(res)
}
