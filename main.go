package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func helloHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.Error(res, "404 not found.", http.StatusNotFound)
		return
	}

	if req.Method != "GET" {
		http.Error(res, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(res, "Hello World!")

}

func main() {
	http.HandleFunc("/", helloHandler)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
