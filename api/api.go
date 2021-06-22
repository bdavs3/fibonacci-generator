package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bdavs3/fibonacci-generator/fib"

	"github.com/gorilla/mux"
)

type Handler struct {
	Generator *fib.Generator
}

func NewHandler(generator *fib.Generator) *Handler {
	return &Handler{
		Generator: generator,
	}
}

func (h *Handler) GetFibonacci(w http.ResponseWriter, r *http.Request) {
	term := mux.Vars(r)["term"]

	intTerm, err := strconv.Atoi(term)
	if err != nil {
		http.Error(w, "term must be an integer", http.StatusBadRequest)
		return
	}

	res := h.Generator.Fibonacci(intTerm)

	fmt.Fprintf(w, "The %dth Fibonacci value is %d.", intTerm, res)
}

func (h *Handler) GetMemoized(w http.ResponseWriter, r *http.Request) {
	val := mux.Vars(r)["val"]

	intVal, err := strconv.Atoi(val)
	if err != nil {
		http.Error(w, "value must be an integer", http.StatusBadRequest)
		return
	}

	res, err := h.Generator.Memoized(intVal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "There are %d memoized terms less than %d.", res, intVal)
}

func (h *Handler) ClearMemoized(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ClearMemoized")
}
