package main

import (
	server "github.com/siangyeh8818/golang.exporter.china.appstore/internal/server"
)

func main() {

	go shellcommand.callpython()
	server.Run_Exporter_Server()
}
