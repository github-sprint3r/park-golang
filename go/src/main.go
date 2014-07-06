package main

import (
	"net/http"
	"fmt"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request) string {
	var requestPath = r.URL.Path[1:]
	fmt.Fprintf(w, "%s", r.URL.Path[1:]) //return url path and print
	return requestPath
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
