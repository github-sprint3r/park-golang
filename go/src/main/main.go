 package main

 import (
	"net/http"
	"parkstore"
)

func main() {
	http.HandleFunc("/parkstore", parkstore.ParkStore)
	http.Handle("/", http.FileServer(http.Dir("../../../public_html/")))
	http.ListenAndServe(":3333", nil)
}
