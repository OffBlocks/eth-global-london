package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	Env                       string `env:"ENV"`
	Port                      string `env:"PORT"`
	DBUrl                     string `env:"DB_URL"`
	ApiBaseUrl                string `env:"API_BASE_URL"`
	TemporalServer            string `env:"TEMPORAL_SERVER"`
	PrivateKey                string `env:"PRIVATE_KEY"`
	AdminAddress              string `env:"ADMIN_ADDRESS"`
	EthereumEndpointUrl       string `env:"ETHEREUM_ENDPOINT_URL"`
	EthereumChainSelector     uint64 `env:"ETHEREUM_CHAIN_SELECTOR"`
	EthereumUsdAddress        string `env:"ETHEREUM_USD_ADDRESS"`
	EthereumUsdDecimals       int    `env:"ETHEREUM_USD_DECIMALS"`
	EthereumEurAddress        string `env:"ETHEREUM_EUR_ADDRESS"`
	EthereumEurDecimals       int    `env:"ETHEREUM_EUR_DECIMALS"`
	EthereumQuoterAddress     string `env:"ETHEREUM_QUOTER_ADDRESS"`
	EthereumSwapRouterAddress string `env:"ETHEREUM_SWAP_ROUTER_ADDRESS"`
	EthereumTransferAddress   string `env:"ETHEREUM_TRANSFER_ADDRESS"`
	BaseEndpointUrl           string `env:"BASE_ENDPOINT_URL"`
	BaseChainSelector         uint64 `env:"BASE_CHAIN_SELECTOR"`
	BaseUsdAddress            string `env:"BASE_USD_ADDRESS"`
	BaseUsdDecimals           int    `env:"BASE_USD_DECIMALS"`
	BaseTransferAddress       string `env:"BASE_TRANSFER_ADDRESS"`
}

func LoadConfig() (config Config, err error) {
	err = godotenv.Load(fmt.Sprintf("./configs/envs/%s.env", os.Getenv("ENV")))
	if err != nil {
		return
	}

	err = env.Parse(&config)
	if err != nil {
		return
	}

	return
}
