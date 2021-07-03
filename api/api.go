package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/bdavs3/fibonacci-generator/db"
	"github.com/bdavs3/fibonacci-generator/fib"
)

const idMatch = "[0-9]+"

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/fib/{term:"+idMatch+"}", GetFibonacci)
	router.HandleFunc("/memoized/{val:"+idMatch+"}", GetMemoized)
	router.HandleFunc("/clear", ClearMemoized)

	return router
}

func GetFibonacci(w http.ResponseWriter, r *http.Request) {
	term := mux.Vars(r)["term"]

	intTerm, err := strconv.Atoi(term)
	if err != nil {
		http.Error(w, "term must be an integer", http.StatusBadRequest)
		return
	}

	res := fib.Fibonacci(intTerm, db.NewConnection())

	fmt.Fprintf(w, "Fibonacci term %d is %d.", intTerm, res)
}

func GetMemoized(w http.ResponseWriter, r *http.Request) {
	val := mux.Vars(r)["val"]

	intVal, err := strconv.Atoi(val)
	if err != nil {
		http.Error(w, "value must be an integer", http.StatusBadRequest)
		return
	}

	res := fib.Memoized(intVal, db.NewConnection())

	fmt.Fprintf(w, "There are %d memoized terms less than %d.", res, intVal)
}

func ClearMemoized(w http.ResponseWriter, r *http.Request) {
	fib.Clear(db.NewConnection())
	fmt.Fprint(w, "Memoized results cleared.")
}
