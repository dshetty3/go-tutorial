package main

import (
	"fmt"
	"time"
	// "strconv"
	"strings"
	"sync"
)

	
	var tourName string= "Post Malone Tour"
	var tourTickets uint = 50
	var remainingTickets uint = 50
	// var bookings = make([]map[string]string, 0)
	var bookings = make([]UserData, 0)


	type UserData struct {
		firstName string
		lastName string
		email string
		numberOfTickets uint
	}

var wg = sync.WaitGroup{}

func main(){

	greetUsers()

	
	// for { 
	firstName, lastName, email, usertickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, usertickets)

	if  isValidName && isValidEmail && isValidTicketNumber {
	bookTickets(remainingTickets, usertickets, firstName, lastName, email)

	wg.Add(1)
	go sendTicket(usertickets, firstName, lastName, email) 


	// for-each
	getFirstNames()
	// fmt.Printf("Booking Names: %v\n", firstNames)
	

	if remainingTickets == 0 {
		fmt.Println("We are sold out!")
		// break
	}
		
	} else {
	// fmt.Printf("We only have %v tickets remaining right now\n", remainingTickets)
	if !isValidName {
		fmt.Println("Pleas enter Valid Name.")
	} 
	if !isValidEmail {
		fmt.Println("Please enter Valid Email ID")
	} 
	if !isValidTicketNumber {
		fmt.Println("Please enter Valid Ticket Number.")
	}
	}
	wg.Wait()
	}

// }

func greetUsers(){
	fmt.Printf("Welcome to the %v Website\n", tourName)
	fmt.Println("Get your Tickets here !")
	fmt.Printf("We currently have %v tickets out of %v tickets !!\n", remainingTickets, tourTickets)
}

func getFirstNames() []string{
	firstNames := []string{}
	for _, booking := range bookings{
		// var names = strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, usertickets uint) (bool, bool, bool){
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := usertickets > 0 && usertickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput()(string, string, string, uint){
	var firstName string
	var lastName string
	var email string
	var usertickets uint

	fmt.Println("Enter your First Name")
	fmt.Scan(&firstName) 

	fmt.Println("Enter Last Name")
	fmt.Scan(&lastName)

	fmt.Println("Enter Email Id")
	fmt.Scan(&email)

	fmt.Println("Enter Number of Tickets needed")
	fmt.Scan(&usertickets)

	return firstName, lastName, email, usertickets

}

func bookTickets(remainingTickets uint, usertickets uint, firstName string, lastName string, email string){
	remainingTickets = remainingTickets - usertickets

	// maps
	// var myslice []string
	// var mymap map[string]string

	// var userData = make(map[string]string)

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: usertickets,
	}
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["emailID"] = email
	// userData["NumberOfTickets"] = strconv.FormatUint(uint64(usertickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of Bookings is %v\n", bookings)

	
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email along with the booking number on %v\n", firstName, lastName, usertickets, email)
	fmt.Printf("We now have %v tickets remaining.\n", remainingTickets)
	// fmt.Printf("Bookings inside function: %v\n", bookings)

}

func sendTicket(usertickets uint, firstName string, lastName string, email string){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", usertickets, firstName, lastName)
	fmt.Println("------------------------")
	fmt.Printf("Sending ticket: %v to email address %v\n", ticket, email)
	fmt.Println("------------------------")
	wg.Done()
}