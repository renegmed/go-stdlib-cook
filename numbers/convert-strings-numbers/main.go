package main

import (
	"fmt"
	"strconv"
)

const bin = "00001"
const hex = "2f"
const intString = "12"
const floatString = "12.3"

func main() {

	// Decimals
	res, err := strconv.Atoi(intString)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\tParsed integer: %d\n", res)

	/*

		The base determines how the number is parsed. Note that the hexadecimal
		has the base (second argument) of 16 and the binary has the base of 2.
		The function Atoi of package strconv is, in fact, the ParseInt function
		with the base of 10.

	*/

	// Parsing hexadecimals
	res64, err := strconv.ParseInt(hex, 16, 32) // 16 - base int,  32 - bitSize int
	if err != nil {
		panic(err)
	}
	fmt.Printf("\tParsed hexadecimal: %d\n", res64)

	// Parsing binary values
	resBin, err := strconv.ParseInt(bin, 2, 32) // 2 - base int,  32 - bitSize int
	if err != nil {
		panic(err)
	}
	fmt.Printf("\tParsed bin: %d\n", resBin)

	// Paseding floating-points
	resFloat, err := strconv.ParseFloat(floatString, 32) // 32 - bitSize int
	if err != nil {
		panic(err)
	}
	fmt.Printf("\tParsed float: %.5f\n", resFloat)

}
