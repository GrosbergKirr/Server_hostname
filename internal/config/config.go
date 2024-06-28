package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	GRPC GRPCConfig
}

type GRPCConfig struct {
	Host    string        `yaml:"host" env-default:"localhost"`
	Port    int           `yaml:"port" env-default:"9090"`
	Timeout time.Duration `yaml:"timeout" env-default:"10h"`
}

func LoadConfig() *Config {
	//Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Get config path from .env
	configPath, exists := os.LookupEnv("CONFIG_PATH")
	if !exists {
		log.Fatal("set CONFIG_PATH env variable")
	}

	// check if file exists
	fmt.Println(configPath)
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file doesn't exist: %s", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("can't read config: %s", err)
	}
	return &cfg
}
