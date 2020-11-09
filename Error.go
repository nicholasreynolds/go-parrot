package main

import "fmt"

type InvalidForwardingAddress struct {
	address string
}

type InvalidPort struct {
	port string
}

func (e InvalidForwardingAddress) Error() string {
	return fmt.Sprintf("invalid forwarding address: %s", e.address)
}

func (e InvalidPort) Error() string {
	return fmt.Sprintf("invalid port: %s", e.port)
}
