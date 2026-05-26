package main 

import "fmt"  

func main() {
	
var conferenceName = "Go Conference"
const conferenceTicket = 50
var remanningTickets = 15

	fmt.Println("Welcome to our", conferenceName,"Booking Application")
	fmt.Println("We have", conferenceTicket, "tickets at the begining, but now only remain", remanningTickets, "tickets")
	fmt.Println("Get your tickets here to attend")
}