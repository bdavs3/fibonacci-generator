package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/bdavs3/fibonacci-generator/fib"
)

func GetFibonacci(w http.ResponseWriter, r *http.Request) {
	term := mux.Vars(r)["term"]

	intTerm, err := strconv.Atoi(term)
	if err != nil {
		http.Error(w, "term must be an integer", http.StatusBadRequest)
		return
	}

	res, err := fib.Fibonacci(intTerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "The %dth Fibonacci value is %d.", intTerm, res)
}

func GetMemoized(w http.ResponseWriter, r *http.Request) {
	val := mux.Vars(r)["val"]

	intVal, err := strconv.Atoi(val)
	if err != nil {
		http.Error(w, "value must be an integer", http.StatusBadRequest)
		return
	}

	res, err := fib.Memoized(intVal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "There are %d memoized terms less than %d.", res, intVal)
}

func ClearMemoized(w http.ResponseWriter, r *http.Request) {
	fib.Clear()
	fmt.Fprint(w, "ClearMemoized")
}
