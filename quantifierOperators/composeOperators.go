package quantifieroperators

// we define the struct that will be used to create the and compose operator
type ComposeAndOperator struct {
	operator1 QuantifierOperator
	operator2 QuantifierOperator
}

// NewComposeOperator is a function that takes two QuantifierOperator values and returns a pointer to a ComposeOperator struct.
func NewComposeOperator(operator1 QuantifierOperator, operator2 QuantifierOperator) *ComposeAndOperator {
	return &ComposeAndOperator{operator1, operator2}
}

// This method returns true if both operators return true, otherwise it returns false.
func (c *ComposeAndOperator) Evaluate() bool {
	return c.operator1.Evaluate() && c.operator2.Evaluate()
}

// we define the struct that will be used to create the compose operator
type ComposeOrOperator struct {
	operator1 QuantifierOperator
	operator2 QuantifierOperator
}

// NewComposeOperator is a function that takes two QuantifierOperator values and returns a pointer to a ComposeOperator struct.
func NewComposeOrOperator(operator1 QuantifierOperator, operator2 QuantifierOperator) *ComposeOrOperator {
	return &ComposeOrOperator{operator1, operator2}
}

// This method returns true if at least one of the operators return true, otherwise it returns false.
func (c *ComposeOrOperator) Evaluate() bool {
	return c.operator1.Evaluate() || c.operator2.Evaluate()
}

// we define the struct that will be used to create the compose operator
type ComposeNotOperator struct {
	operator QuantifierOperator
}

// NewComposeOperator is a function that takes two QuantifierOperator values and returns a pointer to a ComposeOperator struct.
func NewComposeNotOperator(operator QuantifierOperator) *ComposeNotOperator {
	return &ComposeNotOperator{operator}
}

// This method returns true if the operator returns false, otherwise it returns true.
func (c *ComposeNotOperator) Evaluate() bool {
	return !c.operator.Evaluate()
}
