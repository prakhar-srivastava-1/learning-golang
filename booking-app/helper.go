// this module is in the same main package
// but this will contain helper functions
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

// getUserInputs gets all the data required for booking ticket/s
// returns (string, string, string, string, int) - User's first and last names, email, country and number of tickets
func getUserInputs() (string, string, string, string, int) {
	// variables to grab all user inputs
	var firstName string
	var lastName string
	var email string
	var country string
	var userTickets int

	// First name and last name
	fmt.Print("Please enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Please enter your last name: ")
	fmt.Scan(&lastName)

	// email
	fmt.Print("Please enter your email: ")
	fmt.Scan(&email)

	// country
	fmt.Print("Please country of conference (Singapore/London/India): ")
	fmt.Scan(&country)

	// tickets
	fmt.Print("Please enter the number of tickets you want to book: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, country, userTickets
}

// validateUserInputs validates the entries from the user
// firstName (string) - User's first name
// lastName (string) - User's last name
// email (string) - User's email
// returns (string) - Error message; empty string if all good
func validateUserInputs(firstName string, lastName string, email string) string {
	if isValidName(firstName) != 0 {
		return nameErrorMap[isValidName(firstName)]
	}
	if isValidName(lastName) != 0 {
		return nameErrorMap[isValidName(lastName)]
	}
	if isValidEmail(email) != 0 {
		return emailErrorMap[isValidEmail(email)]
	}
	return ""
}

// global constants
const HR = "==============================================================================================="

// greetUsers prints a few welcome messages for the user
// conferenceName (string) - Name of the conference
// remainingTickets (int) - Number of conference tickets remaining
// conferenceTickets (int) - Total number of conference tickets in the beginning
func greetUsers(conferenceName string, remainingTickets int, conferenceTickets uint) {
	fmt.Println(HR)
	fmt.Printf("Welcome to %v Ticket Booking Sytem. ", conferenceName)
	fmt.Printf("We have %d ticket/s remaining out of %d tickets.\n", remainingTickets, conferenceTickets)
	fmt.Println(HR)
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

// bookTickets books tickets for a user
// remainingTickets (int) - Number of tickets remaining
// bookings ([]string) - Slice of all ticket booking transactions
// returns ([]string, int) - Updated slice of bookings and number of remaining tickets
func bookTickets(remainingTickets int, bookings []string) ([]string, int) {
	// get all inputs
	firstName, lastName, email, country, userTickets := getUserInputs()

	// validate form input
	if validateUserInputs(firstName, lastName, email) != "" {
		fmt.Printf("Invalid user input: %s", validateUserInputs(firstName, lastName, email))
		return bookings, remainingTickets
	}

	// check if there are sufficient tickets remaining for this booking
	if remainingTickets < userTickets {
		fmt.Printf("Sorry! We only have %d ticket/s remaining.\n", remainingTickets)
		return bookings, remainingTickets
	}

	// book tickets
	bookings = append(bookings, firstName+" "+lastName)
	// update ticket count
	remainingTickets -= userTickets
	fmt.Printf("Dear %s, you have successfully booked %d tickets for %s conference. Details will be shared with you shortly on %s.\n", firstName, userTickets, country, email)

	return bookings, remainingTickets
}
