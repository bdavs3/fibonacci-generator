package main

import (
	"log"
	"net/http"

	"github.com/bdavs3/fibonacci-generator/api"
	"github.com/bdavs3/fibonacci-generator/fib"

	"github.com/gorilla/mux"
)

func main() {
	generator := fib.NewGenerator()
	handler := api.NewHandler(generator)

	router := mux.NewRouter()

	router.HandleFunc("/fib", handler.GetFibonacci).Methods(http.MethodGet)
	router.HandleFunc("/memoized", handler.GetMemoized).Methods(http.MethodGet)
	router.HandleFunc("/clear", handler.ClearMemoized).Methods(http.MethodDelete)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
