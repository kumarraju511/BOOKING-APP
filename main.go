package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of  %v tickets and %v\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	//ask user for their name
	fmt.Println("Enter user firstname:")
	fmt.Scan(&firstName)

	fmt.Println("Enter user lasttname:")
	fmt.Scan(&lastName)

	fmt.Println("Enter user email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaing for %v", remainingTickets, conferenceName)

}
