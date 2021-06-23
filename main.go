package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/bdavs3/fibonacci-generator/api"
)

const port = "8080"

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/fib/{term}", api.GetFibonacci).Methods(http.MethodGet)
	router.HandleFunc("/memoized/{val}", api.GetMemoized).Methods(http.MethodGet)
	router.HandleFunc("/clear", api.ClearMemoized).Methods(http.MethodDelete)

	fmt.Printf("Listening on %v...", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
