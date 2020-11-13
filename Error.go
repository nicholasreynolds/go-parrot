package main

import "fmt"

type InvalidForwardingAddress struct {
	address string
}

type InvalidPort struct {
	port string
}

type InvalidRoute struct {
	route string
}

func (e InvalidForwardingAddress) Error() string {
	return fmt.Sprintf("invalid forwarding address: %s", e.address)
}

func (e InvalidPort) Error() string {
	return fmt.Sprintf("invalid port: %s", e.port)
}

func (e InvalidRoute) Error() string {
	return fmt.Sprintf("invalid route: %s", e.route)
}
