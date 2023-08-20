package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	// const tells Go that the value is not allowed to change
	const conferenceTickets = 50
	var remainingTickets = 10

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application. ", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

	var userName string
	var userTicket int
	// ask user for their name
	userName = "Tom"
	userTicket = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTicket)

}
