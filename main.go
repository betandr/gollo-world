package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		s := r.URL.Path
		w.Header().Set("Content-Type", "text/html")
		hello(w, s)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
	return
}

func hello(out io.Writer, msg string) {
	if msg == "/" {
		msg = "world"
	} else {
		msg = url.QueryEscape(msg[1:])
	}
	result := fmt.Sprintf("<html>hello, %s!</html>", msg)

	io.WriteString(out, result)
}
