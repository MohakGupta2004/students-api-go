package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Addr string `yaml:"address" env-required:"true"`
}
type Config struct {
	Env        string `yaml:"env" env:"ENV" env-required:"true"`
	Db         string `yaml:"db" env-required:"true"`
	HttpServer `yaml:"http_server"`
}

func MustHave() *Config {
	getEnv := os.Getenv("env")

	if getEnv == "" {
		flags := flag.String("config", "", "path to configuration file")
		flag.Parse()
		getEnv = *flags
		if getEnv == "" {
			log.Fatal("Config file not set")
		}
	}

	if _, conf := os.Stat(getEnv); os.IsNotExist(conf) {
		log.Fatal("configuration file not found")
	}
	var cfg Config
	err := cleanenv.ReadConfig(getEnv, &cfg)
	if err != nil {
		log.Fatalf("Can't able to read from env: %s", err)
	}
	return &cfg
}
