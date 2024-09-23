package main

import (
	"fmt"
	"ws-go/src/server"
)

const (
	HOST = "localhost"
	PORT = "8080"
	CERT_FILE = "server.crt"
	KEY_FILE = "server.key"
)

func main() {
	srv := server.New(HOST, PORT)
	if err := srv.Run(CERT_FILE, KEY_FILE); err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
	fmt.Println("Hello, CompileDaemon!")
}
