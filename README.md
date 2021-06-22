# Fibonacci Generator

### Specifications
Expose a Fibonacci sequence generator through a web API that memoizes intermediate values. The web API should expose operations to (a) fetch the Fibonacci number given an ordinal (e.g. Fib(11) == 89, Fib(12) == 144), (b) fetch the number of memoized results less than a given value (e.g. there are 12 intermediate results less than 120), and (c) clear the data store.

The web API must be written in Go, and Postgres must be used as the data store for the memoized results. Please include tests for your solution, and a README describing how to build and run it.

Bonus points:
- Use dockertest.
- Include a Makefile.
- Include some data on performance.