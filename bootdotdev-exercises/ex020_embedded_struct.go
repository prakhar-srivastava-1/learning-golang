/*
Embedded Structs
Go is not an object-oriented language. However, embedded structs provide a kind of data-only inheritance that can be useful at times. Keep in mind, Go doesn't support classes or inheritance in the complete sense, embedded structs are just a way to elevate and share fields between struct definitions.

	type car struct {
	  make string
	  model string
	}

	type truck struct {
	  // "car" is embedded, so the definition of a
	  // "truck" now also additionally contains all
	  // of the fields of the car struct
	  car
	  bedSize int
	}

Embedded vs nested
An embedded struct's fields are accessed at the top level, unlike nested structs.
Promoted fields can be accessed like normal fields except that they can't be used in composite literals

	lanesTruck := truck{
	  bedSize: 10,
	  car: car{
	    make: "toyota",
	    model: "camry",
	  },
	}

fmt.Println(lanesTruck.bedSize)

// embedded fields promoted to the top-level
// instead of lanesTruck.car.make
fmt.Println(lanesTruck.make)
fmt.Println(lanesTruck.model)
Assignment
At Textio, a "user" struct represents an account holder, and a "sender" is just a "user" with some "sender" specific data. A "sender" is a user that has a rateLimit field that tells us how many messages they are allowed to send.

Fix the system by using an embedded struct as expected by the testEmbeddedStruct code.
*/
package main

import "fmt"

type sender struct {
	rateLimit int
	embeddedUser
}

type embeddedUser struct {
	name   string
	number int
}

// don't edit below this line

func testEmbeddedStruct(s sender) {
	fmt.Println("Sender name:", s.name)
	fmt.Println("Sender number:", s.number)
	fmt.Println("Sender rateLimit:", s.rateLimit)
	fmt.Println("====================================")
}

func embeddedStruct() {
	testEmbeddedStruct(sender{
		rateLimit: 10000,
		embeddedUser: embeddedUser{
			name:   "Deborah",
			number: 18055558790,
		},
	})
	testEmbeddedStruct(sender{
		rateLimit: 5000,
		embeddedUser: embeddedUser{
			name:   "Sarah",
			number: 19055558790,
		},
	})
	testEmbeddedStruct(sender{
		rateLimit: 1000,
		embeddedUser: embeddedUser{
			name:   "Sally",
			number: 19055558790,
		},
	})
}
