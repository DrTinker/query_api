package client

import (
	"query_api/models"
	"sync"
)

type ConfigClient interface {
	Load(path string) error
	GetRPCConfig() (*models.RpcConfig, error)
	GetHttpConfig() (*models.HttpConfig, error)
}

var (
	configClient ConfigClient
	configOnce   sync.Once
)

func GetConfigClient() ConfigClient {
	return configClient
}

func InitConfigClient(client ConfigClient) {
	configOnce.Do(
		func() {
			configClient = client
		},
	)
}
