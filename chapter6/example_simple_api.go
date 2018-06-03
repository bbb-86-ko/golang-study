package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Hello, golang")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
