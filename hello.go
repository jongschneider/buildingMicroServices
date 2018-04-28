package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// *********	types	*********
type helloWorldResponse struct {
	Message string `json:"message"`
}

type hotdog int

// *********	Handlers	*********
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	msg := helloWorldResponse{"Hello World"}
	json.NewEncoder(w).Encode(&msg)
}

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "hotdog!"
	json.NewEncoder(w).Encode(&msg)
}

func main() {
	port := 8080
	var h hotdog
	/*
		The HandleFunc function is a convenience function that creates a handler who's ServeHTTP
		method calls an ordinary function with the func(ResponseWriter, *Request) signature that you pass as a parameter.
	*/
	http.HandleFunc("/helloworld", helloWorldHandler)

	/*
		The Handle function requires that you pass two parameters, the pattern that you would
		like to register the handler and an object that implements the Handler interface:

		type Handler interface {
			ServeHTTP(ResponseWriter, *Request)
		  }
	*/
	// http.Handle("/helloworld", http.HandlerFunc(helloWorldHandler))
	http.Handle("/hotdog", h)

	log.Printf("Now listening on port %v...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
