/*
Same Line Declarations
We are able to declare multiple variables on the same line:

mileage, company := 80276, "Tesla"

// is the same as

mileage := 80276
company := "Tesla"
Assignment
Within the main function, declare a float called averageOpenRate and string called displayMessage on the same line.

Initialize them to values of .23 and is the average open rate of your messages respectively before they are printed.
*/
package main

import "fmt"

func sameLineDeclaration() {
	// solution starts
	// declare here
	averageOpenRate, displayMessage := 0.23, "is the average open rate of your messages"
	// solution ends

	fmt.Println(averageOpenRate, displayMessage)
}
