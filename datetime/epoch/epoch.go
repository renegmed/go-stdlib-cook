package main

import (
	"fmt"
	"time"
)

/*

	The epoch is the universal system to describe the point in time.
	The beginning of epoch time is defined as 00:00:00 1 Jan 1970 UTC.
	The value of epoch is the amount of seconds since the timestamp,
	minus the amount of leap seconds since then.

*/
func main() {

	/*

		To obtain the epoch value from the Time instance, the method with
		the same name as the creation of Time from epoch, Unix, can be called.
		There is one more method called UnixNano, which returns the count of
		milliseconds instead of seconds.

	*/

	// Set the epoch form int64
	t := time.Unix(0, 0)
	fmt.Println(t) // 1969-12-31 19:00:00 -0500 EST

	// Get the epoch
	// from Time instance
	epoch := t.Unix()
	fmt.Println(epoch) // 0

	// Current epoch time
	epochNow := time.Now().Unix()
	fmt.Printf("Epoch time in seconds: %d\n", epochNow) // Epoch time in seconds: 1537916539

	epochNano := time.Now().UnixNano()
	fmt.Printf("Epoch time in nano-seconds: %d\n", epochNano) // 1537916602363483000

}
