package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Conference Tickests is %T, Remainig tickets is %T, Conference name is %T.\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v and %v are available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your shit here:")

	var bookings []string

	for {
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

		if userTickets < int(remainingTickets) {
			fmt.Printf("We only have %v tickets remaining so you can't book %v tickets", remainingTickets, userTickets)
			continue
		}
		remainingTickets = remainingTickets - uint(userTickets)

		bookings = append(bookings, firstName+" "+lastName)
		fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation to %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}

		for _, booking := range bookings {

			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of bookings are %v \n", firstNames)

		noTicketsLeft := remainingTickets == 0
		if noTicketsLeft {
			fmt.Printf("The %v is booked out", conferenceName)
			break
		}
	}

}
