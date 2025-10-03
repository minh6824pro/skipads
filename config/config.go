package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"time"
)

type Config struct {
	ServiceName string `env:"SERVICE_NAME" env-default:"user-skip-ads-service"`
	Http        Http
	MySQL       MySQL
	Redis       Redis
}

type Http struct {
	Port         uint          `env:"HTTP_PORT" env-default:"8080"`
	Timeout      time.Duration `yaml:"Timeout" env-default:"15s"`
	ReadTimeout  time.Duration `yaml:"ReadTimeout" env-default:"5s"`
	WriteTimeout time.Duration `yaml:"WriteTimeout" env-default:"5s"`
}
type MySQL struct {
	URI           string        `env:"MYSQL_URI" env-required:"true"`
	TimeToConnect time.Duration `env:"MYSQL_TIME_TO_CONNECT" env-default:"20s"`

	MaxConnections        uint32 `env:"MYSQL_MAX_CONNECTIONS" env-default:"130"`
	MaxIdleConnections    uint32 `env:"MYSQL_MAX_IDLE_CONNECTIONS" env-default:"10"`
	ConnectionMaxLifeTime uint32 `env:"MYSQL_MAX_CONNECTIONS_TIME" env-default:"30"`
	ConnectionMaxIdleTime uint32 `env:"MYSQL_MAX_IDLE_CONNECTIONS_TIME" env-default:"10"`
}

type Redis struct {
	URI      string `env:"REDIS_URI" env-required:"true"`
	Password string `env:"REDIS_PASSWORD" env-default:""`
	Db       int    `env:"REDIS_DB" env-default:"0"`
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

func (c Config) GetInternalAPIKey() string {
	return InternalAPIKey
}
