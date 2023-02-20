package helper

import "strings"

// ValidateUserInputs - checks if user inputted data is logically valid
// @param fName - first name data to check
// @param lName - last name data to check
// @param email - email data to check
// @param userTickets - requested tickets data to check
// @param remainingTickets - remaining tickets to ensure user does not overbook
// @return (bool, bool, bool) - three booleans, representing whether first and last name,
// 								email, and requested ticket amount are valid, respectively
func ValidateUserInputs(fName string, lName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(fName) >= 2 && len(lName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNum := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNum
}
