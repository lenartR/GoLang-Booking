package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available!\n", conferenceTickets, remainingTickets)
	fmt.Println("Get yout tickets here.")

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
	remainingTickets -= userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a conformation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v remaining tickets for %v\n", remainingTickets, conferenceName)
}
