package handler

import (
	"strconv"
	"strings"
)

type RoundRobin struct{}

func (r RoundRobin) GetRedirectionTarget(nat map[int]string, start *int) (string, int) {
	*start = (*start + 1) % len(nat)
	data := strings.Split(nat[*start], ":")
	host := data[0]
	port, err := strconv.Atoi(data[1])
	if err != nil {
		panic("Invalid port provided.")
	}
	return host, port
}
