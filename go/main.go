 package main

 import (
	"net/http"
	"parkstore"
)

func main() {
	http.HandleFunc("/parkstore", parkstore.ParkStore)
	http.ListenAndServe(":8000", nil)
}
