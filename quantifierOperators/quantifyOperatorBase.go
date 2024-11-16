package quantifieroperators

// we define the interface that all operators must implement
type QuantifierOperator interface {
	// Evaluate is a method that compares a set of fields based on the struct associated and return a boolean value.
	Evaluate() bool
}
