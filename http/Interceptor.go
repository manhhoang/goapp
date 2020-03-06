package main

import (
	"fmt"
	"log"
	"net/http"
)

type a struct{}

func (a a) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Implement server logics here")
}

func main() {
	var h http.Handler = a{}

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Println("Before")
		defer log.Println("After")

		h.ServeHTTP(w, req)
	})

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

