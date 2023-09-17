package pkg

import (
	"github.com/google/wire"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/auth"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/cache"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/config"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/crypto"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

var ProviderSet = wire.NewSet(
	config.ProviderSet,
	cache.ProviderSet,
	wire.Bind(new(repository.Connector), new(repository.DataManager)),
	repository.ProviderSet,
	server.ProviderSet,
	crypto.ProviderSet,
	auth.ProviderSet,
)
