package main

import "log"

func main() {
	// 链接合约
	client, err := ethclient.Dial("wss://eth-sepolia.g.alchemy.com/v2/<API_KEY>")
	if err != nil {
		log.Fatal(err)
	}
}
