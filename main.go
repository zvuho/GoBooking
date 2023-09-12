package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets, city := getUserInput()

	isValidName, isValidEmail, isValidNumTickets := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

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

		bookTicket(userTickets, firstName, lastName, email)
		wg.Add(1) // only one wg needed for send ticket
		go sendTicket(userTickets, firstName, lastName, email)

		var firstNames = getFirstNames()

		fmt.Printf("The first names of bookings are %v \n", firstNames)

		noTicketsLeft := remainingTickets == 0
		if noTicketsLeft {
			fmt.Printf("The %v is booked out", conferenceName)
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
	wg.Wait()
}

func greetUsers() {

	fmt.Printf("Conference Tickests is %T, Remainig tickets is %T, Conference name is %T.\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v and %v are available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your shit here:")

}

func getFirstNames() []string {

	firstNames := []string{}

	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint, string) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	var city string

	fmt.Println("Enter your first name:")

	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")

	fmt.Scan(&lastName)

	fmt.Println("Enter your email name:")

	fmt.Scan(&email)

	fmt.Println("How many tickets do you need:")

	fmt.Scan(&userTickets)

	fmt.Println("The city of the conference:")

	fmt.Scan(&city)

	return firstName, lastName, email, userTickets, city
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - uint(userTickets)
	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation to %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v \n", userTickets, firstName, lastName)
	fmt.Printf("Sending ticket:\n %v\n To email %v\n", ticket, email)
	wg.Done() //finaliza el estado de espera para este thread
}
