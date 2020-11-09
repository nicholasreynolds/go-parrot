package main

import (
	"fmt"
	"net/http"
)

var p Parrot
func main()  {
	var err error
	p, err = GetParrot()
	if err != nil {
		fmt.Errorf(err.Error())
	}
	http.HandleFunc("/", HandleRoot)
	http.ListenAndServe(fmt.Sprintf(":%s", p.port), nil)
}

func HandleRoot(resp http.ResponseWriter, req *http.Request){
	fmt.Println(req)

	handler := Handler{resp, req}
	handler.Handle(req.Method)
}
