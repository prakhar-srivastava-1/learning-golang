// all go code should belong to a package
package main

// imports come at top
// for multiple imports use ()
// each line will contain one packet
import (
	"fmt"
)

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
