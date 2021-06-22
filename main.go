package main

import (
	"log"
	"net/http"

	"github.com/bdavs3/fibonacci-generator/api"
	"github.com/bdavs3/fibonacci-generator/fib"
)

func main() {
	generator := fib.NewGenerator()
	handler := api.NewHandler(generator)

	http.HandleFunc("/fib", handler.GetFibonacci)
	http.HandleFunc("/memoized", handler.GetMemoized)
	http.HandleFunc("/clear", handler.ClearMemoized)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
