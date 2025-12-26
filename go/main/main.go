package main

import (
	"fmt"
	"new_nft_go/config"
	"new_nft_go/database"
	"new_nft_go/eth"
	"new_nft_go/listener"
	"new_nft_go/server"
)

func main() {
	// 链接合约

	cfg := config.InitConfigYaml()
	client := eth.GetClient(cfg)
	go listener.StartListener(client, cfg)
	if client == nil {
		fmt.Println("Failed to connect to Ethereum client")
	}
	fmt.Println("Connected to Ethereum client")
	go database.InitDataBase()
	r := server.StartServer()
	err := r.Run(":8080")
	if err != nil {
		return
	}

}
