package main

import (
	"github.com/asaskevich/govalidator"
	"os"
)

type Parrot struct {
	forwarding string
	port string
}

func GetParrot() (Parrot, error) {
	p := Parrot {
		forwarding: os.Getenv("URL"),
		port: os.Getenv("PORT"),
	}

	if !isURLValid(p.forwarding) {
		return Parrot{}, InvalidForwardingAddress{p.forwarding}
	}
	if !isPortValid(p.port) {
		return Parrot{}, InvalidPort{p.port}
	}
	return p, nil
}

func isURLValid(url string) bool {
	return govalidator.IsURL(url)
}

func isPortValid(port string) bool {
	return govalidator.IsInt(port)
}

