package fib

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func Fibonacci(term int, conn *pgxpool.Pool) int {
	var memoizedVal int
	err := conn.QueryRow(
		context.Background(),
		fmt.Sprintf("SELECT val FROM fib_memo WHERE term = %d;", term),
	).Scan(&memoizedVal)
	if err != nil { // No memoized value found
		if term <= 1 {
			insertMemoized(term, term, conn)
			return term
		}

		result := Fibonacci(term-1, conn) + Fibonacci(term-2, conn)
		insertMemoized(term, result, conn)

		return result
	}

	return memoizedVal
}

func insertMemoized(term, val int, conn *pgxpool.Pool) {
	conn.Query(
		context.Background(),
		fmt.Sprintf("INSERT INTO fib_memo (term, val) VALUES (%d, %d);", term, val),
	)
}

func Memoized(val int, conn *pgxpool.Pool) int {
	var count int
	conn.QueryRow(
		context.Background(),
		fmt.Sprintf("SELECT COUNT(*) FROM fib_memo WHERE val < %d;", val),
	).Scan(&count)

	return count
}

func Clear(conn *pgxpool.Pool) {
	conn.Query(context.Background(), "DELETE FROM fib_memo;")
}
