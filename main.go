package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of  %v tickets and %v\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
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

		if userTickets > remainingTickets {
			fmt.Printf("we only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			continue
		}

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v \n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaing for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of booking are: %v\n", firstNames)

		if remainingTickets == 0 {
			//end program
			fmt.Println("Sorry, conference is sold out")
			break
		}
	}
}
