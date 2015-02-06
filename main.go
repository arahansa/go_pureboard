package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello world Board Server")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello I will be a Board. %s", r.URL.Path[1:])
}
