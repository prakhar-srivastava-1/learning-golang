/*
Conditionals
if statements in Go don't use parentheses around the condition:

	if height > 4 {
	    fmt.Println("You are tall enough!")
	}

else if and else are supported as you would expect:

	if height > 6 {
	    fmt.Println("You are super tall!")
	} else if height > 4 {

	    fmt.Println("You are tall enough!")
	} else {

	    fmt.Println("You are not tall enough!")
	}

Assignment
Fix the bug on line 12. The if statement should print "Message sent" if the messageLen is less than or equal to the maxMessageLen, or "Message not sent" otherwise.

Tips
Here are some of the comparison operators in Go:

== equal to
!= not equal to
< less than
> greater than
<= less than or equal to
>= greater than or equal to
*/
package main

import "fmt"

func conditionals() {
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen, "and a max length of:", maxMessageLen)

	// don't touch above this line
	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}
}
