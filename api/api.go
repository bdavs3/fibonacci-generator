package api

import (
	"net/http"

	"github.com/bdavs3/fibonacci-generator/fib"
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

}

func (h *Handler) GetMemoized(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) ClearMemoized(w http.ResponseWriter, r *http.Request) {

}
