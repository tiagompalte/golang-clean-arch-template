package configs

type ConfigDatabase struct {
	DriverName       DatabaseType `mapstructure:"DRIVER_NAME"`
	ConnectionSource string       `mapstructure:"CONNECTION_SOURCE"`
}

type ConfigRedis struct {
	Host   string `mapstructure:"HOST"`
	Port   int    `mapstructure:"PORT"`
	DB     int    `mapstructure:"DB"`
	Pass   string `mapstructure:"PASS"`
	Prefix string `mapstructure:"PREFIX"`
}

type ConfigCache struct {
	DriverName CacheType   `mapstructure:"DRIVER_NAME"`
	Redis      ConfigRedis `mapstructure:"REDIS"`
}

type ConfigBcrypt struct {
	Round int `mapstructure:"ROUND"`
}

type ConfigJwt struct {
	Algorithm string `mapstructure:"ALGORITHM"`
	KeySecret string `mapstructure:"KEY_SECRET"`
	Duration  int    `mapstructure:"DURATION"`
}

type ConfigMigrate struct {
	DriverName     MigrateType `mapstructure:"DRIVER_NAME"`
	PathMigrations string      `mapstructure:"PATH_MIGRATIONS"`
}

type Config struct {
	AppName  string         `mapstructure:"APP_NAME"`
	WebPort  string         `mapstructure:"WEB_PORT"`
	Database ConfigDatabase `mapstructure:"DATABASE"`
	Cache    ConfigCache    `mapstructure:"CACHE"`
	Bcrypt   ConfigBcrypt   `mapstructure:"BCRYPT"`
	Jwt      ConfigJwt      `mapstructure:"JWT"`
	Migrate  ConfigMigrate  `mapstructure:"MIGRATE"`
}
