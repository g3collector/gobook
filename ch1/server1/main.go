package main

import (
	"fmt"
	"log"
	"net/http"
)

// minimal echo server
func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Url.Path= %q", r.URL.Path)
}
