package main

import (
	"fmt"
	"strings"
)

func greetUsers(conferenceName string, remainingTickets int, conferenceTickets uint) {
	fmt.Printf("Welcome to %v Ticket Booking Sytem.", conferenceName)
	fmt.Printf("We have %d ticket/s remaining out of %d tickets.\n", remainingTickets, conferenceTickets)
}

func isValidName(name string) bool {
	// name must have atleast 2 characters
	// name must not contain any special characters
	if len(name) < 2 {
		return false
	} else if strings.ContainsAny(name, "~`!@#$%^&*()_+-=1234567890|}{[]\\'\";:?/>.<,}") {
		return false
	}
	return true
}

func isValidEmail(email string) bool {
	return strings.Contains(email, "@")
}

func main() {
	const conferenceName = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets = 50

	var firstName string
	var lastName string
	var email string
	var country string
	var userTickets int

	// slice to track bookings
	bookings := []string{}
	// slice to hold first names of all bookings
	firstNames := []string{}

	for {
		greetUsers(conferenceName, remainingTickets, conferenceTickets)

		// First name and last name
		fmt.Print("Please enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Print("Please enter your last name: ")
		fmt.Scan(&lastName)
		if !(isValidName(firstName) && isValidName(lastName)) {
			fmt.Println("First and last name should be at least 2 characters long\nand must not contain and special characters.")
			continue
		}

		// email
		fmt.Print("Please enter your email: ")
		fmt.Scan(&email)
		if !isValidEmail(email) {
			fmt.Println("Invalid email.")
			continue
		}

		// country
		fmt.Print("Please country of conference (Singapore/London/India): ")
		fmt.Scan(&country)

		// tickets
		fmt.Print("Please enter the number of tickets you want to book: ")
		fmt.Scan(&userTickets)

		if remainingTickets < userTickets {
			fmt.Printf("Sorry! We only have %d ticket/s remaining.\n", remainingTickets)
			continue
		}

		// book tickets
		bookings = append(bookings, firstName+" "+lastName)
		remainingTickets -= userTickets
		fmt.Printf("Dear %s, you have successfully booked %d tickets. Details will be shared with you shortly.\n", firstName, userTickets)

		// check if there are any tickets remaining
		if remainingTickets == 0 {
			fmt.Println("We are now out of tickets. See you in the next conference.")
			// print first names of all users who have booked tickets
			for _, booking := range bookings {
				// names gets [firstName, lastName]
				names := strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("People who have booked tickets: %v\n", firstNames)
			break
		}
	}
}
