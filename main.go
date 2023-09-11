package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int
		var city string

		fmt.Println("Enter your first name:")

		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")

		fmt.Scan(&lastName)

		fmt.Println("Enter your email name:")

		isValidName := len(firstName) >= 2 && len(lastName) >= 2

		isValidEmail := strings.Contains(email, "@")

		isValidNumTickets := userTickets >= 0 && userTickets <= int(remainingTickets)

		fmt.Scan(&email)

		fmt.Println("How many tickets do you need:")

		fmt.Scan(&userTickets)

		switch city {
		case "London", "Berlin":
			//exec code
		case "Singapore", "Hong Kong":
			//exec other code
		case "New York":
			//exec code for London conference
		case "Mexico City":
			//exec code for London conference
		default:
			fmt.Println("No valid city selected")
		}

		if isValidName && isValidEmail && isValidNumTickets {

			remainingTickets = remainingTickets - uint(userTickets)

			bookings = append(bookings, firstName+" "+lastName)
			fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation to %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			var firstNames = getFirstNames(bookings)

			fmt.Printf("The first names of bookings are %v \n", firstNames)

			noTicketsLeft := remainingTickets == 0
			if noTicketsLeft {
				fmt.Printf("The %v is booked out", conferenceName)
				break
			}

		} else {
			if !isValidName {
				fmt.Println("Your input data is invalid. Try again.")
			}
			if !isValidEmail {
				fmt.Println("Your input data is invalid. Try again.")
			}
			if !isValidNumTickets {
				fmt.Println("Your input data is invalid. Try again.")
			}

		}
	}

}

func greetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {

	fmt.Printf("Conference Tickests is %T, Remainig tickets is %T, Conference name is %T.\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v and %v are available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your shit here:")

}

func getFirstNames(bookings []string) []string {

	firstNames := []string{}

	for _, booking := range bookings {

		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}
