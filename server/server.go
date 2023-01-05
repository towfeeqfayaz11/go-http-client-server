package main

import (
	"fmt"
	"io"
	"net/http"
)

func rootEnd(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my http website!\n")
}
func helloEnd(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func main() {
	http.HandleFunc("/", rootEnd)
	http.HandleFunc("/hello", helloEnd)
	http.ListenAndServe(":3333", nil)
}
