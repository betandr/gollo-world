package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", hello)
	http.Handle("/", r)
	fmt.Println("Starting on 8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "hello world!")
}
