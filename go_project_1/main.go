package main

import "fmt"

func main() {
	// one can add a colon to a variable, but not a const or a var
	conferenceName := "Go conference"
	// const tells Go that the value is not allowed to change
	const conferenceTickets int = 50
	var remainingTickets int = 10

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application. ", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

	var firstName string
	var lastName string
	var email string
	var userTicket int
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

	userTicket = 2
	fmt.Printf("Thank you, %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTicket, email)

}
