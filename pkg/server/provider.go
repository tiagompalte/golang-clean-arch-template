package server

import (
	"github.com/tiagompalte/golang-clean-arch-template/configs"
)

func ProviderSet(
	config configs.Config,
) Server {
	if config.WebServer == "fiber" {
		return NewFiberServer(config)
	}
	panic("None web server define")
}
