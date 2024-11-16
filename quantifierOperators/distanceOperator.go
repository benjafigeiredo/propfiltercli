package quantifieroperators

// we define the struct that will be used to create the distance operator
// this is struct is defined assuming that the user will find properties in a range of distance from a point
// We can design a more complex distance operator, where we can define a QuantifierOperator interface as a field
// and use it to compare the distance between two points. Like this:

// type DistanceOperator struct {
// 	startPoint        []float64
// 	endPoint          []float64
// 	distanceToCompare float64
// 	operator          QuantifierOperator
// }

type DistanceOperator struct {
	startPoint        []float64
	endPoint          []float64
	distanceToCompare float64
}

// NewDistanceOperator is a function that takes three float64 values and returns a pointer to a DistanceOperator struct.
func NewDistanceOperator(startPoint []float64, endPoint []float64, distanceToCompare float64) *DistanceOperator {
	if len(startPoint) != len(endPoint) {
		panic("The number of dimensions of the start and end points must be the same")
	}

	return &DistanceOperator{startPoint, endPoint, distanceToCompare}
}

// since we are working on real state, it's useful to define a distance operator that takes two points and a distance to compare
// in here, we assume that the user will find properties in a range of distance from where he is. This is a simple implementation
func (d *DistanceOperator) Evaluate() bool {
	// we calculate the distance between the two points
	// we use the Pythagorean theorem to calculate the distance
	// distance = sqrt((x2 - x1)^2 + (y2 - y1)^2)
	// we assume that the points have the same number of dimensions
	var sum float64
	for i := 0; i < len(d.startPoint); i++ {
		sum += (d.endPoint[i] - d.startPoint[i]) * (d.endPoint[i] - d.startPoint[i])
	}

	operator := NewComposeOrOperator(NewLessThanOperator(sum, d.distanceToCompare), NewEqualToOperator(sum, d.distanceToCompare))
	// this will return true if the property is in the range of distance from the point specified.
	return operator.Evaluate()
}
