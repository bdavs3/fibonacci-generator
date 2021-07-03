windows:
	GOOS=windows go build -o bin/fib main.go

mac:
	GOOS=darwin go build -o bin/fib main.go

linux:
	GOOS=linux go build -o bin/fib main.go

test:
	go test ./...