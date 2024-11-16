package quantifieroperators

// we define the struct that will be used to create the include operator
type IncludeOperator struct {
	value1 string
	values []string
}

// NewIncludeOperator is a function that takes a string value and a slice of string values and returns a pointer to an IncludeOperator struct.
func NewIncludeOperator(value1 string, values []string) *IncludeOperator {
	return &IncludeOperator{value1, values}
}

// This method returns true if value1 is included in the values slice, otherwise it returns false.
func (i *IncludeOperator) Evaluate() bool {
	for _, value := range i.values {
		if value == i.value1 {
			return true
		}
	}
	return false
}
