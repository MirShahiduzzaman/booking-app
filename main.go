package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50 // unsigned int - only 0 and positive values
var bookings = make([]UserData, 0)

// custom data type UserData to store user information
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	// welcoming users to application
	greetUsers()

	// retrieving user input data
	firstName, lastName, email, userTickets := getUserInputs()

	// checks if data is valid
	isValidName, isValidEmail, isValidTicketNum := helper.ValidateUserInputs(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNum {
		// if valid input then book tickets
		bookTicket(userTickets, firstName, lastName, email)

		// adds 1 concurrent program to wait group
		wg.Add(1)

		// runs sendTicket command in background
		go sendTicket(userTickets, firstName, lastName, email)

		// print first names of people who booked tickets
		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are : %v\n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Printf("Our conference is booked out! Please come back next year.")
		}
	} else {
		// give details for which input was wrong
		if !isValidName {
			fmt.Printf("First and last names must be at least 2 characters. Please try again.\n")
		}
		if !isValidEmail {
			fmt.Printf("Email address not valid. Please try again.\n")
		}
		if !isValidTicketNum {
			fmt.Printf("Sorry, cannot book %v tickets. Please check the amount.\n", userTickets)
		}
	}

	wg.Wait()
}

var wg = sync.WaitGroup{}

// greetUsers - welcomes the user to the program
func greetUsers() {
	fmt.Printf("Welcome to %s booking application!\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are still available!\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend... :)")
}

// getFirstNames - from the bookings list, retrieves the first names of all ticket holders
// returns []string - first names of all ticket holders
func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		// appends to list of first names
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

// getUserInputs - asks for and retrieves ticket holder data
// @return (string, string, string, uint) - first name, last name, email, and number
//
//	of user requested tickets in that order
func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their first name
	fmt.Print("Enter your first name : ")
	fmt.Scan(&firstName)

	// ask user for their last name
	fmt.Print("Enter your last name : ")
	fmt.Scan(&lastName)

	// ask user for their email address
	fmt.Print("Enter your email address : ")
	fmt.Scan(&email)

	// ask user for requested number of tickets
	fmt.Print("Enter number of tickets : ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

// bookTicket - books the number of tickets that the user specified
// @param userTickets - number of user requested tickets
// @param firstName - first name of ticket holder
// @param lastName - last name of ticket holder
// @param email - email of ticket holder
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// initialize userData var
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

// sendTicket - simulates sending tickets to email in the background
// @param userTickets - number of tickets to send
// @param firstName - first name of ticket holder
// @param lastName - last name of ticket holder
// @param email - email of ticket holder
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("################")
	fmt.Printf("Sending ticket:\n%v \nto email address %v\n", ticket, email)
	fmt.Println("################")

	// removes 1 from the number of concurrent programs we need to wait for
	wg.Done()
}
