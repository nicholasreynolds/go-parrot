package main

import "fmt"

type InvalidForwardingAddress struct {
	address string
}

func (e InvalidForwardingAddress) Error() string {
	return fmt.Sprintf("invalid forwarding address: %s", e.address)
}
