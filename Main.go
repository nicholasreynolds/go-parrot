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
	http.ListenAndServe(":8080", nil)
}

func HandleRoot(resp http.ResponseWriter, req *http.Request){
	fmt.Println(req)

	handler := Handler{resp, req}
	handler.Handle(req.Method)
}
