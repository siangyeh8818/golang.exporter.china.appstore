package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	server "github.com/siangyeh8818/golang.exporter.china.appstore/internal/server"
	shellcommand "github.com/siangyeh8818/golang.exporter.china.appstore/internal/shellcommand"
)

func main() {

	setupSigusr1Trap()

	go func() {
		shellcommand.Callpython()
	}()

	//shellcommand.Callpython()

	server.Run_Exporter_Server()
}

func setupSigusr1Trap() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)
	go func() {
		for range c {
			DumpStacks()
		}
	}()
}
func DumpStacks() {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, true)]
	fmt.Printf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===", buf)
}
