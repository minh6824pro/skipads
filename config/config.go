package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	MySQL MySQL
	Redis Redis
}
type MySQL struct {
	URI                   string `env:"MYSQL_URI" env-required:"true"`
	MaxConnections        uint32 `env:"MYSQL_MAX_CONNECTIONS" env-default:"100"`
	MaxIdleConnections    uint32 `env:"MYSQL_MAX_IDLE_CONNECTIONS" env-default:"10"`
	ConnectionMaxLifeTime uint32 `env:"MYSQL_MAX_CONNECTIONS_TIME" env-default:"30"`
	ConnectionMaxIdleTime uint32 `env:"MYSQL_MAX_IDLE_CONNECTIONS_TIME" env-default:"10"`
}

type Redis struct {
	URI      string `env:"REDIS_URI" env-required:"true"`
	Password string `env:"REDIS_PASSWORD" env-default:""`
	dbString string `env:"REDIS_DB" env-default:"0"`
}

func NewConfig() (Config, error) {

	// load config from env
	var cfg Config
	err := cleanenv.ReadConfig(".env", &cfg)

	//err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
