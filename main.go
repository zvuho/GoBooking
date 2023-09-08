package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Conference Tickests is %T, Remainig tickets is %T, Conference name is %T.\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v and %v are available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your shit here:")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name:")

	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")

	fmt.Scan(&lastName)

	fmt.Println("Enter your email name:")

	fmt.Scan(&email)

	fmt.Println("How many tickets do you need:")

	fmt.Scan(&userTickets)

	fmt.Printf("User %v has %v tickets", userName, userTickets)

}
