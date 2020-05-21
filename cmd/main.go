package main

import (
	server "github.com/siangyeh8818/golang.exporter.china.appstore/internal/server"
	//shellcommand "github.com/siangyeh8818/golang.exporter.china.appstore/internal/shellcommand"
)

func main() {

	//go shellcommand.callpython()
	server.Run_Exporter_Server()
}
