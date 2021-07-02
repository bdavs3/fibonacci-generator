package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

const url = "postgres://localhost:5432/mydatabase"

func NewConnection() *pgxpool.Pool {
	conn, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		fmt.Println(err)
	}

	return conn
}
