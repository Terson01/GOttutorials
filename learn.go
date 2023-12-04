package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Terson conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}
	// var bookings [] string

	fmt.Println("welcome to", conferenceName, "boooking tickets application")
	fmt.Println("we have a total of", conferenceTickets, "tickets and", remainingTickets, "are still availiable to purchase")
	fmt.Println("Get your tickets here")

	fmt.Printf("welcome to %v boooking tickets application\n", conferenceName)
	fmt.Printf("we have a total of %v tickets and %v are still availiable to purchase\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here")

	for {
		var firstname string
		var lastname string
		var email string
		var concertno int
		var userTickets uint

		// we are collecting information from the user to ge their name and amount of tickets purchased.
		// userName= "Peterson"
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstname)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastname)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter your concert identity no: ")
		fmt.Scan(&concertno)

		fmt.Println("Enter the number of Tickets purchased: ")
		fmt.Scan(&userTickets)
		// userTickets= 4

		remainingTickets = remainingTickets - userTickets
		//adding a value in an array
		// bookings[0]=firstname+" "+lastname

		//adding a value into a slice
		bookings = append(bookings, firstname+" "+lastname)

		fmt.Printf("the whole array: %v\n", bookings)
		fmt.Printf("the first value: %v\n", bookings[0])
		fmt.Printf("Array type: %T\n", bookings)
		fmt.Printf("Array length: %v\n", len(bookings))

		fmt.Printf("Thank you %v %v, Your id concert no is %v and you purchased %v Tickets, we would send you an email to %v please endevor to be on time for the concert", firstname, lastname, concertno, userTickets, email)
		fmt.Printf("\n %v tickets id remaining for the %v\n", remainingTickets, conferenceName)

		firstnames := []string{}
		for _, bookers := range bookings {
			names := strings.Fields(bookers)
			firstname := names[0]
			firstnames = append(firstnames, firstname)
		}
		fmt.Printf("list of the first names 0f the bookers %v\n", firstnames)
		fmt.Printf("these are the list of all the ours bookings: %v\n", bookings)

	}

}

// fvjhgjbmjh
