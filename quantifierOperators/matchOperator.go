package quantifieroperators

import "strings"

// we define the struct that will be used to create the match operator
type MatchOperator struct {
	value1 string
	value2 string
}

// NewMatchOperator is a function that takes two string values and returns a pointer to a MatchOperator struct.
func NewMatchOperator(value1 string, value2 string) *MatchOperator {
	return &MatchOperator{value1, value2}
}

// This method returns true if value1 is a substring of value2, otherwise it returns false.
func (m *MatchOperator) Evaluate() bool {
	return strings.Contains(m.value2, m.value1)
}
