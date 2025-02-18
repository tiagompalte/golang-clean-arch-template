package configs

// Database
type DatabaseType string

const (
	DatabaseMySql    DatabaseType = "mysql"
	DatabasePostgres DatabaseType = "postgres"
)

func (t DatabaseType) String() string {
	return string(t)
}

// Cache
type CacheType string

const (
	CacheMemory CacheType = "memory"
	CacheMock   CacheType = "mock"
	CacheRedis  CacheType = "redis"
)

// Migrate
type MigrateType string

const (
	GolangMigrate MigrateType = "golang-migrate"
	NativeMigrate MigrateType = "native"
)
