package main

import (
	"fmt"
	"strings"
)

const conferenceTickets uint = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{} // Ulternative way of creating a slice

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v Tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])

	}
	return firstNames
}
func askUserInfos() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("How many tickets do you want to order?")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketsNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketsNumber
}
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf(" Hi! %v %v, thank you for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets are left for %v\n", remainingTickets, conferenceName)
}

func main() {

	greetUsers()
	//firstNames := []string{}

	//Keep asking for a new booking after one booking is done
	for {

		// Ask the user for their information
		firstName, lastName, email, userTickets := askUserInfos()

		//Validating user input
		isValidName, isValidEmail, isValidTicketsNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidEmail && isValidName && isValidTicketsNumber {
			bookTicket(userTickets, firstName, lastName, email)
			// Add the first name to the slice
			//firstNames := append(firstNames, firstName)

			// Call function printFirtNames
			firstNames := getFirstNames(bookings)
			fmt.Printf("List of first mames who booked tickets are: %v\n", firstNames)

			if remainingTickets == 0 {
				// End the program
				fmt.Println("Our conference is booked out. please come back next year.")
				break
			}
			//fmt.Printf("List of first names who booked tickets are: %v\n", firstNames)

		} else {
			if !isValidName {
				fmt.Println(" Sorry the first Name or last name is too short, Please enter a valid input")
			}

			if !isValidEmail {
				fmt.Println(" Sorry the email entered is invalid, Please enter a valid email")
			}

			if !isValidTicketsNumber {
				fmt.Println(" Sorry the number you entered is invalid, Please enter a valid number")
			}
		}

	}

}
