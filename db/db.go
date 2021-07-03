package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

const url = "postgres://localhost:5432/postgres"

func NewConnection() *pgxpool.Pool {
	conn, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error connecting to db: %s", err.Error()))
	}

	if !TableExists("fib_memo", conn) {
		CreateMemo("fib_memo", conn)
	}

	return conn
}

func TableExists(tbl string, conn *pgxpool.Pool) bool {
	var exists bool
	conn.QueryRow(
		context.Background(),
		fmt.Sprintf(
			`SELECT EXISTS 
			(
				SELECT 1
				FROM information_schema.tables 
				WHERE table_name = '%s'
			);`, tbl),
	).Scan(&exists)

	return exists
}

func CreateMemo(tbl string, conn *pgxpool.Pool) {
	conn.Query(
		context.Background(),
		fmt.Sprintf(
			`CREATE TABLE %s (
				term int,
				val int
			);`, tbl),
	)
}
