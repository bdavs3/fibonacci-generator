package fib

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

const url = "postgres://localhost:5432/mydatabase"

func Fibonacci(term int) (int, error) {
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return -1, err
	}

	var memoizedVal int

	err = conn.QueryRow(
		context.Background(),
		fmt.Sprintf("SELECT val FROM fib_memo WHERE term = %d;", term),
	).Scan(&memoizedVal)
	if err != nil { // No memoized value found
		if term <= 1 {
			insertMemoized(conn, term, term)
			return term, nil
		}

		// Can ignore the next two errors since connection has already been tested at this point
		prev, _ := Fibonacci(term - 1)
		prev2, _ := Fibonacci(term - 2)
		result := prev + prev2
		insertMemoized(conn, term, result)

		return result, nil
	}

	return memoizedVal, nil
}

func insertMemoized(conn *pgx.Conn, term, val int) {
	conn.Query(
		context.Background(),
		fmt.Sprintf("INSERT INTO fib_memo (term, val) VALUES (%d, %d);", term, val),
	)
}

func Memoized(val int) (int, error) {
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return -1, nil
	}

	var count int

	err = conn.QueryRow(
		context.Background(),
		fmt.Sprintf("SELECT COUNT(*) FROM fib_memo WHERE val < %d;", val),
	).Scan(&count)
	if err != nil {
		return -1, err
	}

	return count, nil
}

func Clear() error {
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return err
	}

	conn.Query(context.Background(), "DELETE FROM fib_memo;")

	return nil
}
