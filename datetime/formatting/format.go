package main

import (
	"fmt"
	"time"
)

/*

	Go uses the referential time value Jan 2 15:04:05 2006 MST to define
	the formatting layout. See the code example for padding options.

	The memo for the reference date is that when given in number form,
	it is represented as 1,2,3,4,5,6,-7. The -7 value means that the MST
	time zone is 7 hours behind the UTC.

*/
func main() {

	tTime := time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local)
	// 5 - day
	// 8 - hour
	// 5 - min
	// 2 - seconds
	// 0 - nanoseconds
	// time.Local - location

	// The formatting is done
	// with use of reference value
	// Jan 2, 15:04:05 2006 MST
	fmt.Printf("\ttTime is: %s\n", tTime.Format("2006/1/2"))

	fmt.Printf("\tThe time is: %s\n", tTime.Format("15:04"))

	// The predefined formats could be used
	fmt.Printf("\tThe time is: %s\n", tTime.Format(time.RFC1123))

	// The formatting supports space padding
	// only for days in Go version 1.9.2
	fmt.Printf("\ttTime is %s\n", tTime.Format("2006/1/_2"))

	// The zero padding is done by adding 0
	fmt.Printf("\ttTime is: %s\n", tTime.Format("2006/01/02"))

	// The fraction with leading zeros use 0s
	fmt.Printf("\ttTime is: %s\n", tTime.Format("15:04:05:00"))

	// The fraction without leading zeros use 9s
	fmt.Printf("\ttTime is: %s\n", tTime.Format("15:04:05.999"))

	// Append format appends the formatted time to given buffer
	fmt.Printf("\t%s\n",
		string(tTime.AppendFormat([]byte("The time is up: "), "03:04PM")))

}
