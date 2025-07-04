/*
Named Return Values - Implicit Returns
Even though a function has named return values, we can still explicitly return values if we want to.

func getCoords() (x, y int){
  return x, y // this is explicit
}
Using this explicit pattern we can even overwrite the return values:

func getCoords() (x, y int){
  return 5, 6 // this is explicit, x and y are NOT returned
}
Otherwise, if we want to return the values defined in the function signature we can just use a naked return (blank return):

func getCoords() (x, y int){
  return // implicitly returns x and y
}
Assignment
Fix the function to return the named values implicitly.
*/
package main

import (
	"fmt"
)

func implicitYearsUntilEvents(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return
}

// don't edit below this line

func testImplicitReturns(age int) {
	fmt.Println("Age:", age)
	yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := implicitYearsUntilEvents(age)
	fmt.Println("You are an adult in", yearsUntilAdult, "years")
	fmt.Println("You can drink in", yearsUntilDrinking, "years")
	fmt.Println("You can rent a car in", yearsUntilCarRental, "years")
	fmt.Println("====================================")
}

func implicitNamedReturns() {
	testImplicitReturns(4)
	testImplicitReturns(10)
	testImplicitReturns(22)
	testImplicitReturns(35)
}