/*
Structs in Go
We use structs in Go to represent structured data. It's often convenient to group different types of variables together. For example, if we want to represent a car we could do the following:

	type car struct {
	  Make string
	  Model string
	  Height int
	  Width int
	}

This creates a new struct type called car. All cars have a Make, Model, Height and Width.

In Go, you will often use a struct to represent information that you would have used a dictionary for in Python, or an object literal for in JavaScript.

Assignment
Complete the messageToSend struct definition. It needs two fields:

phoneNumber - an integer
message - a string.
*/
package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message     string
}

// don't edit below this line

func testStructInGo(m messageToSend) {
	fmt.Printf("Sending message: '%s' to: %v\n", m.message, m.phoneNumber)
	fmt.Println("====================================")
}

func structInGo() {
	testStructInGo(messageToSend{
		phoneNumber: 148255510981,
		message:     "Thanks for signing up",
	})
	testStructInGo(messageToSend{
		phoneNumber: 148255510982,
		message:     "Love to have you aboard!",
	})
	testStructInGo(messageToSend{
		phoneNumber: 148255510983,
		message:     "We're so excited to have you",
	})
}
