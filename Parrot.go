package main

import (
	"github.com/asaskevich/govalidator"
	"os"
)

type Parrot struct {
	forwarding string
}

func GetParrot() (Parrot, error) {
	p := Parrot{forwarding: os.Args[1]}
	if !isURLValid(p.forwarding) {
		return Parrot{}, InvalidForwardingAddress{p.forwarding}
	}
	return p, nil
}

func isURLValid(api string) bool {
	return govalidator.IsURL(api)
}

