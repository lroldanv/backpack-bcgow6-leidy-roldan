package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/hackaton-go-bases-main/internal/file"
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/hackaton-go-bases-main/internal/service"
)

func main() {

	var tickets []service.Ticket

	Newfile := file.File{"./tickets.csv"}

	tickets, err := Newfile.ReadTickets()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Funcion para obtener tickets del archivo csv
	newBookings := service.NewBookings(tickets)

	NewTicket, err := service.Bookings.Create(newBookings, "", "pepito@gmail.com", "Spain", "23/04/2013", 10)
	if err != nil {
		fmt.Println(errors.Unwrap(err))
	}

	fmt.Println(NewTicket)

}
