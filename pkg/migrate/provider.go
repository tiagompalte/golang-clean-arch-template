package migrate

import (
	"github.com/tiagompalte/golang-clean-arch-template/configs"
	nativemigrate "github.com/tiagompalte/golang-clean-arch-template/pkg/migrate/native"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
)

func ProviderSet(
	config configs.Config,
	data repository.DataManager,
) Migrate {
	if config.Migrate.Migrate == "golang-migrate" {
		return NewGolangMigrate(config)
	}
	return nativemigrate.NewNativeMigrate(data, config.Migrate)
}
