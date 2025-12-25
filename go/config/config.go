package config

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

var config Config

type Config struct {
	RPCURL                string
	ContractAddressString string
	ContractAddress       common.Address
	MySQLDSN              string
	HTTPPort              string

	// Redis 相關配置
	RedisHost     string
	RedisPort     string
	RedisPassword string // 可選，如果 Redis 設置了密碼
}

func InitConfigYaml() Config {
	viper.SetConfigName("config") // 配置文件名（无扩展名）
	viper.SetConfigType("yaml")   // 如果是JSON，这里会是"json"等
	viper.AddConfigPath(".")      // 查找配置文件的路径，这里表示当前目录
	err := viper.ReadInConfig()   // 查找并读取配置文件
	if err != nil {               // 处理读取配置文件的错误
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	ethClientUrl := viper.GetString("ethClientUrl")
	privateKey := viper.GetString("privateKey")
	contractAddressString := viper.GetString("contractAddress")
	contractAddress := common.HexToAddress(contractAddressString)
	fmt.Println("ethClientUrl:", ethClientUrl)
	fmt.Println("privateKey:", privateKey)

	config = Config{
		RPCURL:                ethClientUrl,
		ContractAddressString: contractAddressString,
		ContractAddress:       contractAddress,
	}
	return config
}
