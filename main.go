package main

import "fmt"

func main() {

	conferenceName := "Go Conference"
	const conferenceTicket = 50
	var remanningTickets uint = 15
	var booking []string

	fmt.Printf("Welcome to our %v Booking Application\n", conferenceName)
	fmt.Println("We have", conferenceTicket, "tickets at the begining, but now only remain", remanningTickets, "tickets")
	fmt.Println("Get your tickets here to attend")

	for {
		var userName string
		var userTickets uint

		fmt.Printf("Enter your name: ")
		fmt.Scan(&userName)

		fmt.Printf("Enter your ticket count: ")
		fmt.Scan(&userTickets)

		remanningTickets = remanningTickets - userTickets
		booking = append(booking, userName)

		fmt.Printf("User %v is booked %v tickets.\n", userName, userTickets)
		fmt.Printf("Remaining Tickets for %v is %v\n", conferenceName, remanningTickets)
		fmt.Printf("These are our all books %v\n", booking)
	}
}
