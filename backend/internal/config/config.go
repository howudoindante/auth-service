package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Jwt struct {
	AccessSecret  string `yaml:"access_secret"  env-required:"true"`
	RefreshSecret string `yaml:"refresh_secret"  env-required:"true"`
}

type Database struct {
	Host               string `yaml:"host"     env-default:"localhost"`
	Port               int    `yaml:"port"     env-default:"5432"`
	User               string `yaml:"user"     env-required:"true"`
	Pass               string `yaml:"pass"     env-required:"true"`
	Name               string `yaml:"name"     env-required:"true"`
	SSLMode            string `yaml:"sslmode"  env-default:"disable"`
	Scheme             string `yaml:"scheme"`
	ConnectingAttempts int    `yaml:"connecting_attempts"  env-default:"5"`
}
type Config struct {
	Env        string     `yaml:"env" env-default:"local"`
	HttpConfig HttpConfig `yaml:"http_server"`
	Database   Database   `yaml:"database"`
	Jwt        Jwt        `yaml:"jwt"`
}

type HttpConfig struct {
	Address     string        `yaml:"address"      env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout"      env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {

	_ = godotenv.Load()

	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read env: %v", err)
	}

	return &cfg
}
