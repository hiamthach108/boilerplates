package adapters

import (
	"boilerplates/configs"

	"github.com/golobby/container/v3"
)

func IocConfigs() {
	container.Singleton(func() *configs.AppConfig {
		appConfig, err := configs.NewAppConfig(".env")
		if err != nil {
			panic(err)
		}
		return appConfig
	})
}
