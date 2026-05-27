package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string        `yaml:"env" env-default:"local"`
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	PG          PGConfig      `yaml:"postgre" env-required:"true"`
	Logging     LoggingConfig `yaml:"logging_handler"`
}

type LoggingConfig struct {
	Level   string `yaml:"level"`
	Format  string `yaml:"format"`
	Discard bool   `yaml:"discard"`
}

type PGConfig struct {
	PORT     int    `yaml:"port"`
	NAME     string `yaml:"name"`
	HOST     string `yaml:"host"`
	PASSWORD string `yaml:"password"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exists: " + path)
	}

	var config Config
	if err := cleanenv.ReadConfig(path, &config); err != nil {
		panic("failed to read config: " + err.Error())
	}
	return &config
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "config/local.yaml", "path to config file")
	flag.Parse()
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res

}
