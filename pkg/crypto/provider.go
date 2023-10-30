package crypto

import "github.com/tiagompalte/golang-clean-arch-template/configs"

func ProviderSet(
	config configs.Config,
) Crypto {
	return NewBcrypt(config.Bcrypt)
}
