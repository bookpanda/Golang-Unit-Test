package main

import (
	"Golang-Unit-Test/leapyear"
	"fmt"
)

func main() {
    for year := 2000; year <= 2100; year++ {
        if leapyear.IsLeapYear(year) {
            fmt.Printf("%d is a leap year\n", year)
        } else {
            fmt.Printf("%d is not a leap year\n", year)
        }
    }
}
