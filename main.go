package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func notImplemented(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not Implemented")
}

func main() {
	fmt.Println("start")
	r := mux.NewRouter()
	r.HandleFunc("/", notImplemented).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
