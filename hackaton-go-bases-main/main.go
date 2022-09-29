package main

import (
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/hackaton-go-bases-main/internal/service"
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/hackaton-go-bases-main/internal/file"
	"fmt"

)

func main() {
	file1 := file.File{"./tickets.csv",}

	fmt.Println(file1.Path)

	fmt.Println(file1.Read())

	var tickets []service.Ticket
	// Funcion para obtener tickets del archivo csv
	service.NewBookings(tickets)
}
