# Fibonacci Generator

### Requirements
- A recent version of Go
- A Postgres server running on your machine, containing the default database 'postgres'

### Usage Instructions
This directory contains a Makefile for building binaries targeted at three different operating systems. They can be built as such:
```sh
$ make windows
$ make linux
$ make mac
```
Tests can also be run using the Makefile:
```sh
$ make test
```
Once the binary has been built, run the server:
```sh
$ bin/fib
Listening on 8080...
```
You may now open a web browser to access the available endpoints:

**localhost:8080/fib/{term}**
Retrieve the nth term of the Fibonacci series.

**localhost:8080/memoized/{val}**
Retrieve the number of memoized values in the Fibonacci series less than the given one.

**localhost:8080/clear**
Clear the cache tracking memoized values.

### Project Specifications
Expose a Fibonacci sequence generator through a web API that memoizes intermediate values. The web API should expose operations to (a) fetch the Fibonacci number given an ordinal (e.g. Fib(11) == 89, Fib(12) == 144), (b) fetch the number of memoized results less than a given value (e.g. there are 12 intermediate results less than 120), and (c) clear the data store.

The web API must be written in Go, and Postgres must be used as the data store for the memoized results. Please include tests for your solution, and a README describing how to build and run it.

Bonus points:
- Use dockertest.
- Include a Makefile.
- Include some data on performance.