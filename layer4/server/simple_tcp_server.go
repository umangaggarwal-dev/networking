package main

import (
	"fmt"

	tcp "github.com/umangsinghal31/networking/tcp"
)

const (
	DEFAULT_HOST = "0.0.0.0"
	DEFAULT_PORT = 8080
)

func main() {
	tcpServer := tcp.NewServer(DEFAULT_HOST, DEFAULT_PORT)
	tcpServer.RegisterHandler(tcpHandler)
	tcpServer.StartListening()
}

func tcpHandler(buf []byte) {
	fmt.Println(string(buf[:]))
}
