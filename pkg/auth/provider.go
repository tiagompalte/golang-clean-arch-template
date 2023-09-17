package auth

import "github.com/tiagompalte/golang-clean-arch-template/configs"

func ProviderSet(
	config configs.Config,
) Auth {
	return NewJwtAuth(config)
}
