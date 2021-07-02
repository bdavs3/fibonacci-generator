package fib

import (
	"testing"

	"github.com/bdavs3/fibonacci-generator/db"
)

const fibMemoTerm = 6

func TestFibonacci(t *testing.T) {
	var tests = []struct {
		comment    string
		term, want int
	}{
		{
			comment: "Fibonacci term 5",
			term:    5,
			want:    5,
		},
		{
			comment: "Fibonacci term 10",
			term:    10,
			want:    55,
		},
		{
			comment: "Fibonacci term 15",
			term:    15,
			want:    610,
		},
		{
			comment: "Positive Fibonacci term 0",
			term:    0,
			want:    0,
		},
		{
			comment: "Positive Fibonacci term 1",
			term:    1,
			want:    1,
		},
	}

	for _, test := range tests {
		t.Run(test.comment, func(t *testing.T) {
			conn := db.NewConnection()

			res := Fibonacci(test.term, conn)

			if res != test.want {
				t.Errorf("got %d, want %d", res, test.want)
			}
		})
	}
}

func TestMemoized(t *testing.T) {
	var tests = []struct {
		comment   string
		val, want int
	}{
		{
			comment: "Memoized terms < 8",
			val:     8,
			want:    6,
		},
		{
			comment: "Memoized terms < 9",
			val:     9,
			want:    7,
		},
		{
			comment: "Memoized terms < 4",
			val:     4,
			want:    5,
		},
		{
			comment: "Memoized terms < 0",
			val:     0,
			want:    0,
		},
	}

	for _, test := range tests {
		t.Run(test.comment, func(t *testing.T) {
			Clear(db.NewConnection())
			Fibonacci(fibMemoTerm, db.NewConnection())

			res := Memoized(test.val, db.NewConnection())
			if res != test.want {
				t.Errorf("got %d, want %d", res, test.want)
			}
		})
	}
}

func TestClear(t *testing.T) {
	t.Run("Clear cache", func(t *testing.T) {
		Fibonacci(fibMemoTerm, db.NewConnection())
		Clear(db.NewConnection())

		res := Memoized(1000, db.NewConnection())
		if res != 0 {
			t.Errorf("got %d, want %d", res, 0)
		}
	})
}
