package matcher

type Matcher struct{}

func NewMatcher() *Matcher {
	return &Matcher{}
}

func (m *Matcher) Match(pattern, path string) bool {
	// Simple exact match implementation
	// Gorilla Mux has more complex pattern matching, this is a basic version
	return pattern == path
}
