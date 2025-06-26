// all go code should belong to a package
package main

// imports come at top
// for multiple imports use ()
// each line will contain one packet
import (
	"fmt"
	"strings"
)

// global constants
const HR = "==============================================================================================="

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
	fmt.Println(HR)
	fmt.Printf("Welcome to %v Ticket Booking Sytem. ", conferenceName)
	fmt.Printf("We have %d ticket/s remaining out of %d tickets.\n", remainingTickets, conferenceTickets)
	fmt.Println(HR)
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

// main function
func main() {
	// constants
	const conferenceName = "Go Conference"
	const conferenceTickets = 50

	// variable to track remaining tickets in any point in time
	var remainingTickets = 50

	// variable to grab user operation choice
	var operationChoice int

	// slice to track bookings
	bookings := []string{}

	// infinite loop to keep booking tickets
	for {
		greetUsers(conferenceName, remainingTickets, conferenceTickets)

		// ask user what does he/she want to do?
		// 1 for booking a new ticket
		// 2 for checking the list of bookings
		// 3 for checking the count of bookings
		// 4 for checking the list of first names
		// 5 for checking the number of remaining tickets
		fmt.Println("Please select one of the following operations:")
		fmt.Println("1 for booking a new ticket")
		fmt.Println("2 for checking the list of bookings")
		fmt.Println("3 for checking the count of bookings")
		fmt.Println("4 for checking the list of first names")
		fmt.Println("5 for checking the number of remaining tickets")
		fmt.Println(HR)
		fmt.Scan(&operationChoice)

		// switch is viable alternative for if-else-if ladder
		switch operationChoice {
		case 1:
			bookings, remainingTickets = bookTickets(remainingTickets, bookings)
		case 2:
			fmt.Printf("List of bookings: %v\n", bookings)
		case 3:
			fmt.Printf("Number of bookings: %d\n", len(bookings))
		case 4:
			fmt.Printf("User first name for bookings: %v\n", printFirstNames(bookings))
		case 5:
			fmt.Printf("Remaining tickets: %d\n", remainingTickets)
		default:
			fmt.Println("Invalid operation. Please retry.")
		}

		// check if there are any tickets remaining
		if remainingTickets == 0 {
			fmt.Println("We are now out of tickets. See you in the next conference.")
			fmt.Printf("People who have booked tickets: %v\n", printFirstNames(bookings))
			break
		}
	}
}
