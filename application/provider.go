package application

import (
	"github.com/google/wire"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/usecase"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/infra"
	"github.com/tiagompalte/golang-clean-arch-template/pkg"
)

var ProviderSet = wire.NewSet(
	pkg.ProviderSet,
	infra.ProviderSet,
	usecase.ProviderSet,
	wire.Struct(new(usecase.UseCase), "*"),
	ProvideApplication,
)
