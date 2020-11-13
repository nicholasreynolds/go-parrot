package main

import (
	"io/ioutil"
	"net/http"
)

type Handler struct {
	resp http.ResponseWriter
	req	*http.Request
	forwarding string
}

func (h Handler) Handle() {
	switch h.req.Method {
	case "OPTIONS":
		h.handleCors()
	default:
		h.handleRequest()
	}
}

func (h Handler) writeCORSHeaders() {
	origin := h.req.Header.Get("Origin")
	acrheaders := h.req.Header.Get("Access-Control-Request-Headers")
	acrmethod := h.req.Header.Get("Access-Control-Request-Method")
	if len(origin) > 0 { h.resp.Header().Set("Access-Control-Allow-Origin", origin) }
	if len(acrheaders) > 0 { h.resp.Header().Set("Access-Control-Allow-Headers", acrheaders) }
	if len(acrmethod) > 0 { h.resp.Header().Set("Access-Control-Allow-Method", acrmethod) }
}

func (h Handler) handleCors() {
	h.writeCORSHeaders()
	h.resp.WriteHeader(200)
}

func (h Handler) handleRequest() {
	// Forward request
	apiResp, err := http.Post(h.forwarding, "text/plain;charset=UTF-8", h.req.Body)
	if err != nil {
		return
	}

	// Write response back to caller
	h.writeCORSHeaders()
	h.resp.WriteHeader(apiResp.StatusCode)
	respBytes, err := ioutil.ReadAll(apiResp.Body)
	apiResp.Body.Close()
	if err != nil {
		return
	}
	h.resp.Write(respBytes)
}
