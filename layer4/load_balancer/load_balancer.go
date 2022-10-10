package main

import (
	"fmt"

	lb_strategy "github.com/umangsinghal31/networking/layer4/load_balancer/strategy"
	tcp "github.com/umangsinghal31/networking/tcp"
)

const (
	DEFAULT_HOST = "0.0.0.0"
	DEFAULT_PORT = 19001
)

var nat map[int]string

type strategy interface {
	GetRedirectionTarget(nat map[int]string, start *int) (string, int)
}

func registerIP() {
	nat = make(map[int]string)
	nat[0] = "0.0.0.0:8000"
	nat[1] = "127.0.0.1:8080"
}

func sendRequest(host string, port int) {
	fmt.Println("Sending request to - host=" + host)
}

func getHandler(s strategy) func(buf []byte) {
	nat_copy := make(map[int]string)
	for key, val := range nat {
		nat_copy[key] = val
	}
	start := 0
	return func(buf []byte) {
		sendRequest(s.GetRedirectionTarget(nat_copy, &start))
	}
}

func main() {
	registerIP()
	tcpServer := tcp.NewServer(DEFAULT_HOST, DEFAULT_PORT)
	handler := getHandler(lb_strategy.RoundRobin{})
	tcpServer.RegisterHandler(handler)
	tcpServer.StartListening()
}
