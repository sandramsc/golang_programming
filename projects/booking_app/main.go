package main

import (
	"fmt"
	"strings"
)

// one can add a colon to a variable, but not a const or a var
var conferenceName = "Go conference"

// const tells Go that the value is not allowed to change
const conferenceTickets int = 50

var remainingTickets uint = 10

// slice
var bookings = []string{}

func main() {
	// call the greetUsers function
	greetUsers()

	// Go only has one type of loop .. the for  loop
	for {

		firstName, lastName, email, userTicket := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTicket)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTicket, firstName, lastName, email)

			// Print first names
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end the program
				fmt.Println("Our conference is booked out, Come back next year.")
				break

			}
		} else {
			if !isValidEmail {
				fmt.Println("first name or last name you entered is too short.")
			}

			if !isValidEmail {
				fmt.Println("email address you entered does not contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets entered is invalid")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	// slice for first names in the form of strings
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)

		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTicket uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTicket > 0 && userTicket <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTicket uint
	// ask user for their name
	// pointer will print out the memory address of the stored variable
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket

}

func bookTicket(userTicket uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTicket
	bookings = append(bookings, firstName+""+lastName)

	//userTicket = 2
	fmt.Printf("Thank you, %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
