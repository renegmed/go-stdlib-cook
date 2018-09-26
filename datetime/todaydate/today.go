package main

import (
	"fmt"
	"time"
)

/*
	The Time type is an instant in time in nanoseconds.
	The zero value of Time is January 1, year 1, 00:00:00.000000000 UTC.

	The pointer to the Time type should not be used. If only the value
	(not pointer to variable) is used, the Time instance is considered
	to be safe for use across multiple goroutines. The only exception
	is with serialization.
*/

func main() {
	today := time.Now()
	fmt.Println(today) // 2018-09-25 18:03:58.53325 -0400 EDT m=+0.000543530

}
