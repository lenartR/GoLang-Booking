package main

//belongs to main package -> it knows where the remaining ticket is!

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName = len(firstName) >= 2 && len(lastName) >= 2
	var isEmailValid = strings.Contains(email, "@")
	var isTicketNumberValid = userTickets > 0 && userTickets < remainingTickets

	return isValidName, isEmailValid, isTicketNumberValid
}
