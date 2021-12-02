package main

import (
	"log"
	"trace_demo/api/server"
	"trace_demo/pkg/common"
)

func main() {
	common.ProjectName = "trace_demo"
	srv := server.Server{}
	err := srv.Startup(":8080")
	log.Fatal(err)
}
