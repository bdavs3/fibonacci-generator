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
		http.Error(w, "Value must be an integer.", http.StatusBadRequest)
		return
	}

	res := h.Generator.Fibonacci(intTerm)

	fmt.Fprintf(w, "The %dth Fibonacci value is %d.", intTerm, res)
}

func (h *Handler) GetMemoized(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetMemoized")
}

func (h *Handler) ClearMemoized(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ClearMemoized")
}
