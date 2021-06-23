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

	var val int

	err = conn.QueryRow(
		context.Background(),
		fmt.Sprintf("SELECT val FROM fib_memo WHERE term = %d;", term),
	).Scan(&val)
	if err != nil { // No memoized value found
		prev, _ := Fibonacci(term - 1)
		prev2, _ := Fibonacci(term - 2)
		result := prev + prev2
		conn.Query(
			context.Background(),
			fmt.Sprintf("INSERT INTO fib_memo (term, val) VALUES (%d, %d);", term, result),
		)

		return result, nil
	}

	return val, nil
}

func Memoized(val int) (int, error) {
	var count int

	conn, _ := pgx.Connect(context.Background(), url)
	err := conn.QueryRow(
		context.Background(),
		fmt.Sprintf("SELECT COUNT(*) FROM fib_memo WHERE val < %d;", val),
	).Scan(&count)
	if err != nil {
		return -1, err
	}

	return count, nil
}

func Clear() {

}
