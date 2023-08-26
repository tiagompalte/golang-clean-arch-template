package repository

import (
	"github.com/tiagompalte/golang-clean-arch-template/configs"
)

func ProviderSet(
	config configs.Config,
) DataManager {
	return NewSqlConnect(config.Database)
}
