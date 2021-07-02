package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/bdavs3/fibonacci-generator/db"
	"github.com/bdavs3/fibonacci-generator/fib"
)

const idMatch = "[0-9]+"

type Handler struct {
	Conn *pgxpool.Pool
}

func NewHandler(conn *pgxpool.Pool) *Handler {
	return &Handler{
		Conn: conn,
	}
}

func Router() *mux.Router {
	conn := db.NewConnection()
	handler := NewHandler(conn)
	router := mux.NewRouter()

	router.HandleFunc("/fib/{term:"+idMatch+"}", handler.GetFibonacci)
	router.HandleFunc("/memoized/{val:"+idMatch+"}", handler.GetMemoized)
	router.HandleFunc("/clear", handler.ClearMemoized)

	return router
}

func (h *Handler) GetFibonacci(w http.ResponseWriter, r *http.Request) {
	term := mux.Vars(r)["term"]

	intTerm, err := strconv.Atoi(term)
	if err != nil {
		http.Error(w, "term must be an integer", http.StatusBadRequest)
		return
	}

	res := fib.Fibonacci(intTerm, h.Conn)

	fmt.Fprintf(w, "Fibonacci term %d is %d.", intTerm, res)
}

func (h *Handler) GetMemoized(w http.ResponseWriter, r *http.Request) {
	val := mux.Vars(r)["val"]

	intVal, err := strconv.Atoi(val)
	if err != nil {
		http.Error(w, "value must be an integer", http.StatusBadRequest)
		return
	}

	res := fib.Memoized(intVal, h.Conn)

	fmt.Fprintf(w, "There are %d memoized terms less than %d.", res, intVal)
}

func (h *Handler) ClearMemoized(w http.ResponseWriter, r *http.Request) {
	fib.Clear(h.Conn)
	fmt.Fprint(w, "Memoized results cleared.")
}
