package main

import (
	"fmt"
	"ws-go/src/server"
)

const (
	HOST = "localhost"
	PORT = "8080"
	CERT_FILE = "/etc/letsencrypt/live/dinoraceroyale.zapto.org/fullchain.pem"
	KEY_FILE = "/etc/letsencrypt/live/dinoraceroyale.zapto.org/privkey.pem"
)

func main() {
	srv := server.New(HOST, PORT)
	if err := srv.Run(CERT_FILE, KEY_FILE); err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
	fmt.Println("Hello, CompileDaemon!")
}
