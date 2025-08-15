package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings []string

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isEmailValid, isTicketNumberValid := validateUserInput(firstName, lastName, email, userTickets)

		if isEmailValid && isValidName && isTicketNumberValid {
			bookTicket(userTickets, firstName, lastName, email)
			//call funciton print first names:
			firstNames := getFirstNames()
			fmt.Printf("These are all the first names of bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				//end program
				fmt.Println("Our conference is sold sold out.")
				break
			}

		} else {
			if !isTicketNumberValid {
				fmt.Printf("We only have %v tickets remaining, so you cant book %v tickets!\n", remainingTickets, userTickets)
			}
			if !isEmailValid {
				fmt.Printf("The email %v is invalid. Plese enter valid email!\n", email)
			}
			if !isValidName {
				fmt.Printf("First name or last name to short. Please try again!\n")
			}
		}
	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available!\n", conferenceTickets, remainingTickets)
	fmt.Println("Get yout tickets here.")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames

}
func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	var isValidName = len(firstName) >= 2 && len(lastName) >= 2
	var isEmailValid = strings.Contains(email, "@")
	var isTicketNumberValid = userTickets > 0 && userTickets < remainingTickets

	return isValidName, isEmailValid, isTicketNumberValid
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email adress: ")
	fmt.Scan(&email)
	fmt.Print("Enter the number of tickets you want: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a conformation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v remaining tickets for %v\n", remainingTickets, conferenceName)

}

//!ARRAY
// var bookings [50]string
// var bookings = [50]string{}
//array size 50 with strings. You can add values in {}
//dodajaš na index bookings[0] = x

// !SLICE is better ne rabš povedat kok dolg bo array
// var slice= [50]string{}
//var bookings []string
//adding value via append(slice, x)
