package repository

import (
	"github.com/tiagompalte/golang-clean-arch-template/configs"
)

func ProviderDataSqlManagerSet(
	config configs.Config,
) DataSqlManager {
	return NewDataSqlWithConfig(config.Database)
}

func ProviderConnectorSqlSet(
	config configs.Config,
) ConnectorSql {
	dataSql := ProviderDataSqlManagerSet(config)
	return dataSql.Command()
}
