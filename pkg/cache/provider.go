package cache

import (
	"github.com/tiagompalte/golang-clean-arch-template/configs"
)

func ProviderSet(
	config configs.Config,
) Cache {
	if config.Cache.DriverName == "memory" {
		return NewMemoryCache()
	} else if config.Cache.DriverName == "redis" {
		return NewRedisCache(
			config.Cache.Redis.Host,
			config.Cache.Redis.Port,
			config.Cache.Redis.DB,
			config.Cache.Redis.Pass,
			config.Cache.Redis.Prefix,
		)
	}
	panic("None cache define")
}
