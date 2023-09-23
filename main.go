package main

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	bookings := []string{} // Ulternative way of creating a slice

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v Tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	//Keep asking for a new booking after one booking is done
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// Ask the user for their information

		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email")
		fmt.Scan(&email)

		fmt.Println("How many tickets do you want to order?")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketsNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidEmail && isValidName && isValidTicketsNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf(" Hi! %v %v, thank you for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

			fmt.Printf("%v tickets are left for %v\n", remainingTickets, conferenceName)

			// Extracting first name from bookings slice
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of bookings are: %v\n", firstNames)

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
}
