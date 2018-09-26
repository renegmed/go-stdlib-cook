package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Date(2017, 11, 29, 21, 1, 2, 3, time.Local)
	fmt.Printf("Extracting units from: %v\n", t) //  2017-11-29 21:00:00 -0500 EST

	dOfMonth := t.Day()
	weekDay := t.Weekday()
	month := t.Month()
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()
	// nsecond := t.Nanosecond()

	fmt.Printf("The %dth day of %v is %v hour %d minute %d second %d\n",
		dOfMonth, month, weekDay, hour, minute, second)
	// The 29th day of November is Wednesday hour 21 minute 1 second 2

}
