/*
Basic Variables
bool: a boolean value, either true or false
string: a sequence of characters
int: a signed integer
float64: a floating-point number
byte: exactly what it sounds like: 8 bits of data
Declaring a Variable the Sad Way
var mySkillIssues int
mySkillIssues = 42

The first line, var mySkillIssues int, defaults the mySkillIssues variable to its zero value, 0. On the next line, 42 overwrites the zero value.

We'll talk about a better way to declare variables in the next lesson.

Assignment
Initialize the variables from the print statement to int, float64, bool and string with their zero values, respectively.
*/

package main

import "fmt"

func basicVariables() {
	// solution begins
	// initialize variables here
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	smsSendingLimit = 0
	costPerSMS = 0.0
	hasPermission = false
	username = ""
	// solution ends

	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
