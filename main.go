package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bdavs3/fibonacci-generator/api"
	"github.com/bdavs3/fibonacci-generator/fib"

	"github.com/gorilla/mux"
)

const port = "8080"

func main() {
	generator, err := fib.NewGenerator()
	if err != nil {
		log.Fatal(err)
	}
	handler := api.NewHandler(generator)

	router := mux.NewRouter()

	router.HandleFunc("/fib/{term}", handler.GetFibonacci).Methods(http.MethodGet)
	router.HandleFunc("/memoized/{val}", handler.GetMemoized).Methods(http.MethodGet)
	router.HandleFunc("/clear", handler.ClearMemoized).Methods(http.MethodDelete)

	fmt.Printf("Listening on %v...", port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
