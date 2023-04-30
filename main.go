package main

import (
	writeToFile "booking/helpers"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type reservation struct {
	name   string
	amount uint32
	date   time.Time
}

func main() {
	conferenceName := "Go Conference"

	const conferenceTickets uint32 = 50
	var remainingTickets uint32 = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("We have a total of %d tickets. There are still %d tickets remaining.\n", conferenceTickets, remainingTickets)

	reservations := []reservation{}
	for {
		// Ask user for input
		var userName string
		var userTickets uint32
		fmt.Print("Enter your username: ")
		fmt.Scan(&userName)
		for {
			fmt.Print("Enter the amount of tickets you want to buy: ")
			fmt.Scan(&userTickets)
			if userTickets <= remainingTickets {
				break
			} else {
				fmt.Println("Invalid input, try again.")
			}
		}

		fmt.Printf("User %s booked %d tickets.\n", userName, userTickets)

		reservations = addReservation(reservations, userName, userTickets)

		remainingTickets -= userTickets
		if remainingTickets > conferenceTickets {
			log.Fatal("There are more tickets than in the beginning.")
		}

		fmt.Printf("There are %d tickets remaining.\n", remainingTickets)

		// end the program
		if remainingTickets <= 0 {

			// convert reservations to string
			var s []string
			for _, v := range reservations {
				s = append(s, v.name, strconv.Itoa(int(v.amount)), v.date.String())
			}

			writeToFile.WriteToFile(string(strings.Join(s, "\n")))
			fmt.Println("----------------------------")
			fmt.Println("All tickets have been sold out!")
			fmt.Printf("ALL RESERVATIONS: %v\n", reservations)
			break
		}
	}

}

func addReservation(arr []reservation, name string, ticketNumber uint32) []reservation {
	var resData reservation = reservation{name, ticketNumber, time.Now()}
	arr = append(arr, resData)
	return arr
}
