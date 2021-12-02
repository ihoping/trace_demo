package configs

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"path/filepath"
)

type RedisConfig struct {
	Name     string
	Host     string
	Port     int
	Password string
	DB       int
}

var RedisTraceDemo RedisConfig

func init() {
	path, err := filepath.Abs("configs/redis/trace_demo.yml")
	if err != nil {
		log.Fatal(err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(data, &RedisTraceDemo)
	if err != nil {
		log.Fatal(err)
	}
}
