// all go code should belong to a package
package main

// imports come at top
// for multiple imports use ()
// each line will contain one packet
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

// greetUsers prints a few welcome messages for the user
// conferenceName (string) - Name of the conference
// remainingTickets (int) - Number of conference tickets remaining
// conferenceTickets (int) - Total number of conference tickets in the beginning
func greetUsers(conferenceName string, remainingTickets int, conferenceTickets uint) {
	fmt.Printf("Welcome to %v Ticket Booking Sytem. ", conferenceName)
	fmt.Printf("We have %d ticket/s remaining out of %d tickets.\n", remainingTickets, conferenceTickets)
}

// isValidName checks if the name string is valid
// name (string) - Name input from user
// returns int - Error code; 0 if no errors
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

// isValidEmail checks if the email string is valid
// email (string) - Email input from user
// returns int - Error code; 0 if no errors
func isValidEmail(email string) int {
	if !strings.Contains(email, "@") {
		return 1
	}
	return 0
}

// printFirstNames returns the first names of all users who booked tickets
// bookings ([]string) - All ticket bookings
// returns ([]string) - First names of all users
func printFirstNames(bookings []string) []string {
	// slice to hold first names of all bookings
	firstNames := []string{}
	// print first names of all users who have booked tickets
	// '_' is used to ignore any variable
	for _, booking := range bookings {
		// names gets [firstName, lastName]
		names := strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

// main function
func main() {
	// constants
	const conferenceName = "Go Conference"
	const conferenceTickets uint = 50

	// variable to track remaining tickets in any point in time
	var remainingTickets = 50

	// variables to grab user inputs
	var firstName string
	var lastName string
	var email string
	var country string
	var userTickets int

	// slice to track bookings
	bookings := []string{}

	// infinite loop to keep booking tickets
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
		// update ticket count
		remainingTickets -= userTickets
		fmt.Printf("Dear %s, you have successfully booked %d tickets. Details will be shared with you shortly.\n", firstName, userTickets)

		// check if there are any tickets remaining
		if remainingTickets == 0 {
			fmt.Println("We are now out of tickets. See you in the next conference.")
			fmt.Printf("People who have booked tickets: %v\n", printFirstNames(bookings))
			break
		}
	}
}
