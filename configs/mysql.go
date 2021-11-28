package configs

import (
	"log"
	"os"
	"path/filepath"
)

type MysqlDSN struct {
	Name string
	DSN  string
}

var MySqlTraceDemo MysqlDSN

func init() {
	path, err := filepath.Abs("configs/mysqldsn/trace_demo")
	if err != nil {
		log.Fatal(err)
	}
	dsnBytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	MySqlTraceDemo = MysqlDSN{
		Name: "trace_demo",
		DSN:  string(dsnBytes),
	}
}
