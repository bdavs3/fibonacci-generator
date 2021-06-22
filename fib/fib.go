package fib

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) Fibonacci(val int) int {
	return -1
}

func (g *Generator) Memoized(val int) int {
	return -1
}

func (g *Generator) Clear() {

}
