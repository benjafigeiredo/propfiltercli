package quantifieroperators

// BaseCompareOperator is a base struct that holds common values for comparison.
type BaseCompareOperator struct {
	value1 float64
	value2 float64
}

// LessThanOperator struct embeds BaseCompareOperator and implements its own Evaluate method.
type LessThanOperator struct {
	BaseCompareOperator
}

// NewLessThanOperator creates a new LessThanOperator.
func NewLessThanOperator(value1, value2 float64) *LessThanOperator {
	return &LessThanOperator{
		BaseCompareOperator: BaseCompareOperator{value1: value1, value2: value2},
	}
}

// Evaluate returns true if value1 is less than value2, otherwise false.
func (l *LessThanOperator) Evaluate() bool {
	return l.value1 < l.value2
}

// GreaterThanOperator struct embeds BaseCompareOperator and implements its own Evaluate method.
type GreaterThanOperator struct {
	BaseCompareOperator
}

// NewGreaterThanOperator creates a new GreaterThanOperator.
func NewGreaterThanOperator(value1, value2 float64) *GreaterThanOperator {
	return &GreaterThanOperator{
		BaseCompareOperator: BaseCompareOperator{value1: value1, value2: value2},
	}
}

// Evaluate returns true if value1 is greater than value2, otherwise false.
func (g *GreaterThanOperator) Evaluate() bool {
	return g.value1 > g.value2
}

// EqualToOperator struct embeds BaseCompareOperator and implements its own Evaluate method.
type EqualToOperator struct {
	BaseCompareOperator
}

// NewEqualToOperator creates a new EqualToOperator.
func NewEqualToOperator(value1, value2 float64) *EqualToOperator {
	return &EqualToOperator{
		BaseCompareOperator: BaseCompareOperator{value1: value1, value2: value2},
	}
}

// Evaluate returns true if value1 is equal to value2, otherwise false.
func (e *EqualToOperator) Evaluate() bool {
	return e.value1 == e.value2
}
