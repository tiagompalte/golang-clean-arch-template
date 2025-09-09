package infra

import (
	"github.com/google/wire"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/infra/mongo"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/infra/sql"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/infra/uow"
)

var ProviderSet = wire.NewSet(
	sql.ProviderSet,
	uow.ProviderSet,
	mongo.ProviderSet,
)
