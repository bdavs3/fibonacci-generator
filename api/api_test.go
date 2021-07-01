package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIRequest(t *testing.T) {
	var tests = []struct {
		comment, endpoint string
		want              int
	}{
		{
			comment:  "Positive Fibonacci term",
			endpoint: "/fib/10",
			want:     http.StatusOK,
		},
		{
			comment:  "Negative Fibonacci term",
			endpoint: "/fib/-5",
			want:     http.StatusNotFound,
		},
		{
			comment:  "Non-integer Fibonacci term",
			endpoint: "/fib/dog",
			want:     http.StatusNotFound,
		},
		{
			comment:  "Positive memoized term",
			endpoint: "/memoized/10",
			want:     http.StatusOK,
		},
		{
			comment:  "Negative memoized term",
			endpoint: "/memoized/-5",
			want:     http.StatusNotFound,
		},
		{
			comment:  "Non-integer memoized term",
			endpoint: "/memoized/dog",
			want:     http.StatusNotFound,
		},
		{
			comment:  "Clear cache",
			endpoint: "/clear",
			want:     http.StatusOK,
		},
	}

	for _, test := range tests {
		rec := httptest.NewRecorder()

		t.Run(test.comment, func(t *testing.T) {
			req, err := http.NewRequest(
				http.MethodGet,
				test.endpoint,
				nil,
			)
			if err != nil {
				t.Errorf("Error forming request: %v", err)
			}

			Router().ServeHTTP(rec, req)

			if rec.Code != test.want {
				t.Errorf("got %d, want %d", rec.Code, test.want)
			}
		})
	}
}
