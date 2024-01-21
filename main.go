package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remaining tickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	var bookings []string

	for remainingTickets > 0 && len(bookings) < 50 {
		var userTickets uint
		var firstName string
		var lastName string
		var email string
		// ask user for their name
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		if userTickets < remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("the whole array: %v\n", bookings)
			fmt.Printf("the first value: %v\n", bookings[0])
			fmt.Printf("Slice type: %T\n", bookings)
			fmt.Printf("Slice length: %v\n", len(bookings))

			fmt.Printf("Thanks you %v %v for booking %v tickets. You will receive a confirmtion email at %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
			fmt.Printf("There are all our bookings: %v\n", bookings)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				var firstName = names[0]
				firstNames = append(firstNames, firstName)
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. come back next year!")
				break
			}

		} else if userTickets == remainingTickets {
			//do something else
			fmt.Print("You have booked exactly the same tickets as the remaining tickets")
		} else {
			fmt.Printf("we only have %v tickets remaining, so you can't book %v tickets.\n", remainingTickets, userTickets)
		}

	}

}
