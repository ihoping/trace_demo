package main

import "trace_demo/api/server"

func main() {
	_ = server.Run(":8080")
	//fmt.Println(filepath.Abs("app"))
}
