package cache

import (
	"github.com/tiagompalte/golang-clean-arch-template/configs"
)

func ProviderSet(
	config configs.Config,
) Cache {
	if config.Cache.DriverName == "memory" {
		return NewMemoryCache()
	}
	panic("None cache define")
}
