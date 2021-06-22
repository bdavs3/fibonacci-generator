package fib

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

const url = "postgres://localhost:5432/mydatabase"

type Generator struct {
	Cache *pgx.Conn
}

func NewGenerator() (*Generator, error) {
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}

	return &Generator{
		Cache: conn,
	}, nil
}

func (g *Generator) Fibonacci(val int) int {
	if val <= 1 {
		return val
	}
	return g.Fibonacci(val-1) + g.Fibonacci(val-2)
}

func (g *Generator) Memoized(val int) (int, error) {
	var count int

	err := g.Cache.QueryRow(
		context.Background(),
		fmt.Sprintf("SELECT COUNT(*) FROM fib_memo WHERE val < %d;", val),
	).Scan(&count)
	if err != nil {
		return -1, err
	}

	return count, nil
}

func (g *Generator) Clear() {

}
