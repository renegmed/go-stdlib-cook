package main

/*

The package math/big offers the Float type that could be configured
for a given precision. The advantage of this package is that the
precision could be much higher than the precision of the float64 type.
For illustrative purposes, the small precision values were used
to show the rounding and comparison in the given precision.

*/
import (
	"fmt"
	"math/big"
)

/*
	Note that the da and db numbers are equal when using the precision
	of 16-bits and not equal when using the precision of 32-bits. The
	maximal configurable precision can be obtained from the big.MaxPrec
	constant.
*/
var da float64 = 0.299999992
var db float64 = 0.299999991

var prec uint = 32
var prec2 uint = 16

func main() {
	fmt.Printf("Comparing float64 with '==' equals: %v\n", da == db)

	daB := big.NewFloat(da).SetPrec(prec)
	dbB := big.NewFloat(db).SetPrec(prec)

	fmt.Printf("\tA: %v\n", daB)
	fmt.Printf("\tB: %v\n", dbB)
	fmt.Printf("\tComparing big.Float with precision: %d : %v\n",
		prec, daB.Cmp(dbB) == 0)

	daB = big.NewFloat(da).SetPrec(prec2)
	dbB = big.NewFloat(db).SetPrec(prec2)

	fmt.Printf("\tA: %v\n", daB)
	fmt.Printf("\tB: %v\n", dbB)
	fmt.Printf("\tComparing big.Float with precision: %d : %v\n",
		prec2, daB.Cmp(dbB) == 0)

}
