package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var INDEX_HTML []byte

func main() {
	fmt.Println("Starting server on http://localhost:3003")
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":3003", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(INDEX_HTML)
}

func init() {
	INDEX_HTML, _ = ioutil.ReadFile("./template/index.html")
}
