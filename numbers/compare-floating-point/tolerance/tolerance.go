package main

import (
	"fmt"
	"math"
)

const da = 0.29999999999999998889776975374843459576368331909180
const db = 0.30

func main() {

	// fmt.Printf("\tNumber db=%.60f\n", db)
	fmt.Printf("\tNumber db=%.70f\n", db)
	fmt.Printf("\tNumber da=%.70f\n", da)
	daStr := fmt.Sprintf("%.10f", da)
	dbStr := fmt.Sprintf("%.10f", db)

	fmt.Printf("\tStrings %s = %s equals: %v \n", daStr, dbStr, dbStr == daStr)

	fmt.Printf("\tNumber db=%.60f da=%.60f are equal?  %v\n", db, da, db == da)

	// As the precision of float representation
	// is limited. For the float compoarison it is
	// better to use comparison with some tolerance.
	fmt.Printf("\tNumber equals with TOLERANCE: %v\n", equals(da, db))
}

const TOLERANCE = 1e-8

// Equals compares the floating-point numbers
// with tolerance 1e-8
func equals(numA, numB float64) bool {
	delta := math.Abs(numA - numB)
	if delta < TOLERANCE {
		return true
	}
	return false
}
