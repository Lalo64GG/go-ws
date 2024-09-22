package main

import (
	"fmt"
	"ws-go/src/server"
)


const (
	HOST = "localhost"
	PORT = "8080"
)

func main(){	
	srv := server.New(HOST, PORT)
	srv.Run()
	fmt.Println("Hello, CompileDaemon!")
}