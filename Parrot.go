package main

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"net/http"
)

type Parrot struct {
	forwarding string
	port string
	route string
}

func GetParrot(forwarding string, port string, route string) (Parrot, error) {
	if !govalidator.IsURL(forwarding) {
		return Parrot{}, InvalidForwardingAddress{forwarding}
	}
	if !govalidator.IsPort(port) {
		return Parrot{}, InvalidPort{port}
	}
	if !isRouteValid(route) {
		return Parrot{}, InvalidRoute{route}
	}
	return Parrot{
		forwarding: forwarding,
		port: port,
		route: route,
	}, nil
}

func isRouteValid(route string) bool {
	return govalidator.StringMatches("(/|[a-z]|[A-Z]|[0-9])+", route)
}

func (p Parrot) ListenAndRepeat() error {
	http.HandleFunc(p.route, p.forward)
	return http.ListenAndServe(fmt.Sprintf(":%s", p.port), nil)
}

// Pass request to handler
func (p Parrot) forward(resp http.ResponseWriter, req *http.Request){
	fmt.Println(*req)
	h := Handler{resp, req, p.forwarding}
	h.Handle()
}
