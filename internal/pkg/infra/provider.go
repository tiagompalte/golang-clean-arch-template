package infra

import (
	"github.com/google/wire"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/infra/data"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/infra/uow"
)

var ProviderSet = wire.NewSet(
	data.ProviderSet,
	uow.ProviderSet,
)
