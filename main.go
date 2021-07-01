package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bdavs3/fibonacci-generator/api"
)

const port = "8080"

func main() {
	router := api.Router()

	fmt.Printf("Listening on %v...", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
