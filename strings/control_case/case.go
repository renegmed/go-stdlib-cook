package main

import (
	"fmt"
	"strings"
	"unicode"
)

const email = "ExampPle@domain.com"
const name = "isaac newton"
const upc = "upc"
const i = "i"

const snakeCase = "first_name"

func main() {

	// For comparing the user input
	// sometimes it is better to
	// compare the input in the same case.
	input := "Example@domain.com"
	input = strings.ToLower(input)
	emailToCompare := strings.ToLower(email)
	matches := input == emailToCompare
	fmt.Printf("\tEmail matches: %t\n", matches)

	upcCode := strings.ToUpper(upc)
	fmt.Printf("\tUPPER case: %s\n", upcCode)

	// This digraph has different upper case and
	// title case.
	str := "ǳ"
	fmt.Printf("\t%s in upper: %s and title: %s \n", str,
		strings.ToUpper(str), strings.ToTitle(str))

	// Use of XXXSpecial function
	title := strings.ToTitle(i)
	titleTurk := strings.ToTitleSpecial(unicode.TurkishCase, i)
	if title != titleTurk {
		fmt.Printf("\tToTitle is defferent: %#U vs. %#U \n",
			title[0], []rune(titleTurk)[0])
	}

	// In some cases the input
	// needs to be corrected in case.
	correctNameCase := strings.Title(name)
	fmt.Printf("\tCorrected name: %s\n", correctNameCase)

	// Converting the snake case
	// to camel case with use of
	// Title and ToLower functions.
	firstNameCamel := toCamelCase(snakeCase)
	fmt.Printf("\tCamel case: %s\n", firstNameCamel)
}

func toCamelCase(input string) string {
	titleSpace := strings.Title(strings.Replace(input, "_", " ", -1))
	camel := strings.Replace(titleSpace, " ", "", -1)
	return strings.ToLower(camel[:1]) + camel[1:]
}
