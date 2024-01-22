package config

type Config struct {
	ZapConfig
	DatabaseConfig
}

type ZapConfig struct {
	Filename   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
}

type DatabaseConfig struct {
	MysqlConfig
	RedisConfig
}
type MysqlConfig struct {
	Addr     string
	Port     string
	DB       string
	Username string
	Password string
}

type RedisConfig struct {
}
