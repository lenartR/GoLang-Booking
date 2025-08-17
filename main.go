package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

//dodaš katere threade čakaš preden se zakluč program.

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isEmailValid, isTicketNumberValid := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isEmailValid && isValidName && isTicketNumberValid {

			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			//če daš spredej go ti nrdi ta thread usporeden /concurent task
			go sendTicket(userTickets, firstName, lastName, email)

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
	//waits for all the added threads to be finished before the main thread finishes
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available!\n", conferenceTickets, remainingTickets)
	fmt.Println("Get yout tickets here.")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames

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

	//create a map for user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookins is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a conformation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v remaining tickets for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(3 * time.Second)
	//generates a ticket
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("----------")
	fmt.Printf("Sending ticket: \n%v \nto email adress %v\n", ticket, email)
	fmt.Println("----------")

	//rempves the thread that we added from the waiting list
	wg.Done()
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
