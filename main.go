package main

import (
	"Booking_app/helper"
	"fmt"
	"strconv"
)

const conferenceTickets uint = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0) //Bookings, is where we store user info and we and we start by initialising a list of maps

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v Tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
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
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//Creating a map for our users
	var userData = make(map[string]string) //This is an empty map
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf(" Hi! %v %v, thank you for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets are left for %v\n", remainingTickets, conferenceName)
}
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])

	}
	return firstNames
}

func main() {

	greetUsers()

	//Keep asking for a new booking after one booking is done
	for {

		// Ask the user for their information
		firstName, lastName, email, userTickets := askUserInfos()

		//Validating user input
		isValidName, isValidEmail, isValidTicketsNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTicketsNumber {
			bookTicket(userTickets, firstName, lastName, email)

			// Call function printFirtNames
			//firstNames := getFirstNames(bookings)
			//fmt.Printf("List of first mames who booked tickets are: %v\n", firstNames)

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
