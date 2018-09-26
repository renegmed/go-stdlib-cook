package main

import (
	"fmt"
	"math"
)

var valA float64 = 3.55554444

func main() {

	// Bad assumption on rounding
	// the number by casting it to
	// integer.
	intVal := int(valA)
	fmt.Printf("\tBad rounding by casting to in: %v\n", intVal)

	/*

		The proper way of rounding the float number is to use the function
		that would also consider the decimal part. The commonly used method
		of rounding is to half away from zero (also known as commercial rounding).
		Put simply, if the number contains the absolute value of the decimal
		part which is greater or equal to 0.5, the number is rounded up, otherwise,
		it is rounded down.

		In the function Round, the function Trunc of package math truncates the
		decimal part of the number. Then, the decimal part of the number is extracted.
		If the value exceeds the limit of 0.5 than the value of 1 with the same sign
		as the integer value is added.

		Go version 1.10 uses a faster implementation of the function mentioned in
		the example. In version 1.10, you can just call the math.Round function to
		get the rounded number.

	*/
	fRound := Round(valA)
	fmt.Printf("\tRounding by custom function: %v\n", fRound)

	fRound = math.Round(valA)
	fmt.Printf("\tRounding by math.Round() function: %v\n", fRound)
}

// Round returns the nearest integer.
func Round(x float64) float64 {
	t := math.Trunc(x)
	if math.Abs(x-t) >= 0.5 {
		return t + math.Copysign(1, x)
	}
	return t
}
