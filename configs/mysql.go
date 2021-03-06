package configs

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"path/filepath"
)

type MysqlConfig struct {
	Name     string
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

var MySqlTraceDemo MysqlConfig

func init() {
	path, err := filepath.Abs("configs/mysql/trace_demo.yml")
	if err != nil {
		log.Fatal(err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(data, &MySqlTraceDemo)
	if err != nil {
		log.Fatal(err)
	}
}
