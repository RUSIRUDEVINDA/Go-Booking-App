package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go Conference"
	const conferenceTicket = 50
	var remanningTickets uint = 15
	var bookings []string

	fmt.Printf("Welcome to our %v Booking Application\n", conferenceName)
	fmt.Println("We have", conferenceTicket, "tickets at the begining, but now only remain", remanningTickets, "tickets")
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var userTickets uint

		fmt.Printf("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Printf("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Printf("Enter your ticket count: ")
		fmt.Scan(&userTickets)

		remanningTickets = remanningTickets - userTickets
		var fullName = firstName + " " + lastName
		bookings = append(bookings, fullName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Split(booking, " ")
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The first names of the bookings are %v\n", firstNames)
		fmt.Printf("Remaining ticket count is %v\n", remanningTickets)
	}
}
