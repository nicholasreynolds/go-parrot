package main

import (
	"fmt"
	"os"
)

func main() {
	printError := func(err error) {
		if err != nil {
			fmt.Errorf(err.Error())
		}
	}
	polly, err := GetParrot(
		os.Getenv("URL"),
		os.Getenv("PORT"),
		os.Getenv("ROUTE"),
	)
	printError(err)
	printError(polly.ListenAndRepeat())
}
