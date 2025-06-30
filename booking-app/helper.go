// this module is in the same main package
// but this will contain helper functions
package main

import (
	"fmt"
	"strings"
	"time"
)

// global constants
const (
	HR        = "==============================================================================================="
	TICKET_HR = "###########################################################"
)

// create a struct to hold user data
// type - creates a custom data type
// UserData - name of the custom data type
// struct - since this is a collection of attributes
// informally, strut ~ class
type UserData struct {
	firstName       string
	lastName        string
	email           string
	venue           string
	numberOfTickets uint8
}

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
// returns (string, string, string, string, int) - User's first and last names, email, venue and number of tickets
func getUserInputs() (string, string, string, string, int) {
	// variables to grab all user inputs
	var firstName string
	var lastName string
	var email string
	var venue string
	var userTickets int

	// First name and last name
	fmt.Print("Please enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Please enter your last name: ")
	fmt.Scan(&lastName)

	// email
	fmt.Print("Please enter your email: ")
	fmt.Scan(&email)

	// venue
	fmt.Print("Please enter venue of conference (Singapore/London/Bengaluru): ")
	fmt.Scan(&venue)

	// tickets
	fmt.Print("Please enter the number of tickets you want to book: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, venue, userTickets
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
func printFirstNames(bookings []UserData) []string {
	// slice to hold first names of all bookings
	firstNames := []string{}
	// print first names of all users who have booked tickets
	// '_' is used to ignore any variable
	for _, booking := range bookings {
		// names gets [firstName, lastName]
		// firstName := booking["firstName"]
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

// bookTickets books tickets for a user
// remainingTickets (int) - Number of tickets remaining
// bookings ([]string) - Slice of all ticket booking transactions
// returns ([]string, int) - Updated slice of bookings and number of remaining tickets
func bookTickets(remainingTickets int, bookings []UserData) ([]UserData, int) {
	// get all inputs
	firstName, lastName, email, venue, userTickets := getUserInputs()

	// create a map for storing all user details
	// alternate syntax - var transaction = make(map[string]string)
	// var transaction = map[string]string{
	// 	"firstName": firstName,
	// 	"lastName":  lastName,
	// 	"email":     email,
	// 	"venue":     venue,
	// 	"tickets":   strconv.FormatInt(int64(userTickets), 10),
	// }

	// instead of map, create use UserData struct
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		venue:           venue,
		numberOfTickets: uint8(userTickets),
	}

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
	bookings = append(bookings, userData)
	// update ticket count
	remainingTickets -= userTickets
	fmt.Printf("Dear %s, you have successfully booked %d tickets for %s conference. Details will be shared with you shortly on %s.\n", firstName, userTickets, venue, email)
	// this function simulates email sending
	// this is blocking code
	// to make it in non-blocking mode
	// we run it as a go-routine
	// add a wait group to ensure main function waits for this thread to complete
	wg.Add(1)
	go sendTicket(uint(userTickets), firstName, lastName, email)

	return bookings, remainingTickets
}

// sendTicket simulates sending of the tickets to user
// userTickets (uint) - Number of tickets booked
// firstName (string) - First name of the user
// lastName (string) - Last name of the user
// email (string) - Email of the user
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticketMessage = fmt.Sprintf("%d tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println(TICKET_HR)
	fmt.Printf("Sending tickets:\n%v to email address %v\n", ticketMessage, email)
	fmt.Println(TICKET_HR)
	// go-routine has completed so decrease the go-routine count in wait group
	wg.Done()
}
