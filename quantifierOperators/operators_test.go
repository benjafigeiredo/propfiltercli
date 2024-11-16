package quantifieroperators

import "testing"

// TestGreaterThanOperator_Evaluate is a function that tests the Evaluate method of the GreaterThanOperator struct.
func TestGreaterThanOperator_Evaluate(t *testing.T) {
	// we define the fields struct that will be used to create the greater than operator
	type fields struct {
		value1 float64
		value2 float64
	}

	// we define the struct to hold test cases, and we create instances
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"Test 1", fields{1.2, 2.4}, false},
		{"Test 2", fields{2.3, 1.1}, true},
		{"Test 3", fields{1.0, 1.0}, false},
	}

	// iterate over the tests and run them
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGreaterThanOperator(tt.fields.value1, tt.fields.value2)
			if got := g.Evaluate(); got != tt.want {
				t.Errorf("GreaterThanOperator.Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestLessThanOperator_Evaluate is a function that tests the Evaluate method of the LessThanOperator struct.
func TestLessThanOperator_Evaluate(t *testing.T) {
	// we define the fields struct that will be used to create the less than operator
	type fields struct {
		value1 float64
		value2 float64
	}

	// we define the struct to hold test cases, and we create instances
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"Test 1", fields{1, 2}, true},
		{"Test 2", fields{2, 1}, false},
		{"Test 3", fields{1, 1}, false},
	}

	// iterate over the tests and run them
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLessThanOperator(tt.fields.value1, tt.fields.value2)
			if got := l.Evaluate(); got != tt.want {
				t.Errorf("LessThanOperator.Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestEqualToOperator_Evaluate is a function that tests the Evaluate method of the EqualToOperator struct.
func TestEqualToOperator_Evaluate(t *testing.T) {
	// we define the fields struct that will be used to create the equal to operator
	type fields struct {
		value1 float64
		value2 float64
	}

	// we define the struct to hold test cases, and we create instances
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"Test 1", fields{1, 2}, false},
		{"Test 2", fields{2, 1}, false},
		{"Test 3", fields{1, 1}, true},
	}

	// iterate over the tests and run them
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewEqualToOperator(tt.fields.value1, tt.fields.value2)
			if got := e.Evaluate(); got != tt.want {
				t.Errorf("EqualToOperator.Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
