package main

import (
	"fmt"
	"strings"
)

func greetUsers(confName string, confTicket uint, remainingTickets uint) {
	fmt.Printf("Welcome to our %v booking application\n", confName)
	fmt.Printf("We have a total of %v Tickets and %v are still available\n", confTicket, remainingTickets)
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
func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketsNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketsNumber
}

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	bookings := []string{} // Ulternative way of creating a slice

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	//Keep asking for a new booking after one booking is done
	for {

		// Ask the user for their information
		firstName, lastName, email, userTickets := askUserInfos()

		//Validating user input
		isValidName, isValidEmail, isValidTicketsNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTicketsNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf(" Hi! %v %v, thank you for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

			fmt.Printf("%v tickets are left for %v\n", remainingTickets, conferenceName)

			// Call function printFirtNames
			firstNames := getFirstNames(bookings)
			fmt.Printf("List of first mames who booked tickets are: %v\n", firstNames)

			if remainingTickets == 0 {
				// End the program
				fmt.Println("Our conference is booked out. please come back next year.")
				break
			}

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
