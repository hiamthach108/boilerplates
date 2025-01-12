package adapters

import (
	"boilerplates/configs"
	"boilerplates/libs/cache"
	"boilerplates/shared/interfaces"
	"log"

	"github.com/golobby/container/v3"
)

func IoCCache() {
	container.SingletonLazy(func() interfaces.ICache {
		var (
			logger    interfaces.ILogger
			appConfig *configs.AppConfig
			err       error
		)

		err = container.Resolve(&logger)
		if err != nil {
			log.Fatal(err)
		}
		err = container.Resolve(&appConfig)
		if err != nil {
			log.Fatal(err)
		}

		return cache.NewAppCache(appConfig, logger)
	})
}
