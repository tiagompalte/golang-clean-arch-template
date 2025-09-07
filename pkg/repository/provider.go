package repository

import (
	"github.com/tiagompalte/golang-clean-arch-template/configs"
)

func ProviderDataSqlManagerSet(
	config configs.Config,
) DataSqlManager {
	return NewDataSqlWithConfig(config.DatabaseSQL)
}

func ProviderConnectorSqlSet(
	config configs.Config,
) ConnectorSql {
	dataSql := ProviderDataSqlManagerSet(config)
	return dataSql.Command()
}

func ProviderDataMongoManagerSet(
	config configs.Config,
) DataMongoManager {
	return NewDataMongoWithConfig(config.DatabaseMongo)
}

func ProviderConnectorMongoSet(
	config configs.Config,
) ConnectorMongo {
	dataMongo := ProviderDataMongoManagerSet(config)
	return dataMongo.Command()
}
