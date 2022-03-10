package start

import (
	"query_api/client"
	"query_api/conf"
	"query_api/infrastructure/config"
)

func initConfig() {
	impl := config.NewConfigClientImpl()
	err := impl.Load(conf.App)
	if err != nil {
		panic(err)
	}

	client.InitConfigClient(impl)
}
