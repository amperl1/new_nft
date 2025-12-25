package eth

import (
	"fmt"
	"log"
	"sync"

	"new_nft_go/config"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	Client     *ethclient.Client
	clientOnce sync.Once
	clientMu   sync.RWMutex
)

func GetClient(config config.Config) *ethclient.Client {
	if Client != nil {
		return Client
	}

	var initErr error

	clientOnce.Do(func() {
		Client, initErr = ethclient.Dial(config.RPCURL)
		if initErr != nil {
			return
		}
		log.Println("✅ Connected to Ethereum node")
	})

	if initErr != nil {
		fmt.Println(initErr)
		return nil
	}

	// 检查连接是否仍然有效
	//clientMu.RLock()
	//defer clientMu.RUnlock()

	if Client == nil {
		fmt.Println("client is not initialized")
		return nil
	}
	fmt.Println("client is initialized")
	return Client
}
