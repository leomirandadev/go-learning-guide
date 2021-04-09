package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Result struct {
	Msg string
}

func main() {
	serveMuxAPI()
	// serveNativeAPI()
}

func serveMuxAPI() {
	router := mux.NewRouter()

	router.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Result{Msg: "Hello World"})
	}).Methods("GET")

	fmt.Println("Serving MUX on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func serveNativeAPI() {
	http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(Result{Msg: "Hello World"})
	})

	fmt.Println("Serving NATIVE HTTP on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
