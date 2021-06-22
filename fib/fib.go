package fib

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) Fibonacci(val int) int {
	if val <= 1 {
		return val
	}
	return g.Fibonacci(val-1) + g.Fibonacci(val-2)
}

func (g *Generator) Memoized(val int) int {
	return -1
}

func (g *Generator) Clear() {

}
