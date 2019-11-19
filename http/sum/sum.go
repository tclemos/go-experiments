package sum

// Summer represents a implementation that knows how to sum
type Summer interface {
	Sum(int, int) int
}

// IntSummer know how to sum two integers
type IntSummer struct{}

func (s *IntSummer) Sum(a int, b int) int {
	return a + b
}
