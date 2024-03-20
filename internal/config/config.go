package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

// Config - struct for storing config
type Config struct {
	Env            string `yaml:"env" env-default:"local" env-required:"true"`
	DatabaseName   string `yaml:"database_name" env-requited:"true"`
	DatabasePath   string `yaml:"database_path" env-required:"true"`
	MigrationsPath string `yaml:"migrations_path" env-required:"true"`
	HTTPServer     `yaml:"http_server"`
}

// HTTPServer - struct for storing HTTP server config details
type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"http://localhost:8888"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

// MustLoad loads config from config file
func MustLoad(configPath string) *Config {

	// check if config path is not null
	if configPath == "" {
		log.Fatal("config path is not set")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	// check if we can read config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
