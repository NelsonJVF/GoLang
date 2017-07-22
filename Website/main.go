package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server on http://localhost:3003")
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":3003", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello."))
}
