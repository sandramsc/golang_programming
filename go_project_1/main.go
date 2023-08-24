package main

import (
	"fmt"
	"strings"
)

func main() {
	// one can add a colon to a variable, but not a const or a var
	conferenceName := "Go conference"
	// const tells Go that the value is not allowed to change
	const conferenceTickets int = 50
	var remainingTickets uint = 10
	// slice
	bookings := []string{}

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application. ", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

	// Go only has one type of loop .. the for  loop
	for {

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

		fmt.Println("Enter your email name: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTicket)

		remainingTickets = remainingTickets - userTicket
		// array
		//bookings[0] = firstName + " " + lastName
		// dynamic lists using slices
		bookings = append(bookings, firstName+""+lastName)

		//userTicket = 2
		fmt.Printf("Thank you, %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTicket, email)
		fmt.Printf("%v tickets remaining for %v. \n", remainingTickets, conferenceName)

		// slice for first names in the form of strings
		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)

			firstNames = append(firstNames, names[0])
		}
		//fmt.Printf("These are all our bookings: %v\n", booking)
		fmt.Printf("The first names of bookings are: %v.\n", firstNames)
	}
}
