package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/rickar/cal/v2"
)

func computePTOHours(numDays int, daysTaken int, rateHoursPerDay float64) time.Duration {
	return time.Duration(float64(numDays-daysTaken)*rateHoursPerDay) * time.Hour
}

func main() {
	start := flag.String(
		"startDate", "", "The date from which accumulation should begin, inclusive. Format as YYYY-MM-DD.",
	)
	daysPerYear := flag.Float64(
		"dpy", 0, "The rate at which you accumulate PTO, in days (8 hours) per year.",
	)
	hoursPerHour := flag.Float64(
		"hph", 0, "The rate at which you accumulate PTO, in hours per hour.",
	)
	flag.Parse()

	startDate, err := time.ParseInLocation("2006-01-02", *start, time.Local)
	if err != nil {
		fmt.Println("-startDate must be formatted as YYYY-MM-DD")
		os.Exit(1)
	}
	fmt.Printf("You started work on %v\n", startDate)

	if *daysPerYear != 0 && *hoursPerHour != 0 {
		fmt.Println("cannot set both -dpy and -hph")
		os.Exit(1)
	}
	var hoursPerDay float64
	if *daysPerYear != 0 {
		// This conversion assumes 261 workdays in a year, which is pretty standard.
		// Doesn't use the `cal` package because that only offers workdays in a
		// specific year, not in a generalized average year.
		hoursPerDay = *daysPerYear * 8 / 261
	} else if *hoursPerHour != 0 {
		hoursPerDay = (*hoursPerHour) * 8
	} else {
		fmt.Println("must supply nonzero value for either -dpy or -hph")
		os.Exit(1)
	}
	fmt.Printf("You earn %f hours of PTO per week that you work.\n", hoursPerDay*5)

	c := cal.NewBusinessCalendar()
	numDays := c.WorkdaysInRange(startDate, time.Now())
	fmt.Printf("You have had %d days of work.\n", numDays)

	pto := computePTOHours(numDays, 0, hoursPerDay)
	fmt.Printf("You have accumulated %.2f hours (%.1f days) of PTO.\n", pto.Hours(), pto.Hours()/8)
}
