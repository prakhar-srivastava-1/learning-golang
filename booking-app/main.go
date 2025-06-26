package main

import (
	"fmt"
	"strings"
)

// a map to hold error codes and messages for names and emails
var nameErrorMap = map[int]string{
	0: "",
	1: "Name should at least be 2 characters",
	2: "Name must not contain any special characters",
}

var emailErrorMap = map[int]string{
	0: "",
	1: "Email must contain '@' character",
}

func greetUsers(conferenceName string, remainingTickets int, conferenceTickets uint) {
	fmt.Printf("Welcome to %v Ticket Booking Sytem. ", conferenceName)
	fmt.Printf("We have %d ticket/s remaining out of %d tickets.\n", remainingTickets, conferenceTickets)
}

func isValidName(name string) int {
	// name must have atleast 2 characters
	// name must not contain any special characters
	if len(name) < 2 {
		return 1
	} else if strings.ContainsAny(name, "~`!@#$%^&*()_+-=1234567890|}{[]\\'\";:?/>.<,}") {
		return 2
	}
	return 0
}

func isValidEmail(email string) int {
	if !strings.Contains(email, "@") {
		return 1
	}
	return 0
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
		if isValidName(firstName) != 0 {
			fmt.Printf("Invalid first name. %s\n", nameErrorMap[isValidName(firstName)])
			continue
		}
		fmt.Print("Please enter your last name: ")
		fmt.Scan(&lastName)
		if isValidName(lastName) != 0 {
			fmt.Printf("Invalid last name. %s\n", nameErrorMap[isValidName(lastName)])
			continue
		}

		// email
		fmt.Print("Please enter your email: ")
		fmt.Scan(&email)
		if isValidEmail(email) != 0 {
			fmt.Printf("Invalid email. %s\n", emailErrorMap[isValidEmail(email)])
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
