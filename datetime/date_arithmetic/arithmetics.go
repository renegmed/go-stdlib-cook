package main

import (
	"fmt"
	"time"
)

func main() {

	loc, err := time.LoadLocation("Europe/Vienna")
	if err != nil {
		panic(err)
	}

	t := time.Date(2017, 11, 30, 11, 10, 20, 0, loc)
	fmt.Printf("\tDefault date is: %v\n", t)
	// Default date is: 2017-11-30 11:10:20 +0100 CET

	// Add 3 days
	r1 := t.Add(72 * time.Hour)
	fmt.Printf("\tDefault date -3HRS is: %v\n", r1)
	// Default date -3HRS is: 2017-12-03 11:10:20 +0100 CET

	// More comfortable api
	// to add days/months/years
	r1 = t.AddDate(1, 3, 2)
	fmt.Printf("\tDefault date +1YR +3MTH + 2D is: %v\n", r1)
	// Default date +1YR +3MTH + 2D is: 2019-03-04 11:10:20 +0100 CET
}
