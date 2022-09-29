package file

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/hackaton-go-bases-main/internal/service"
)

type File struct {
	Path string
}

func (f *File) ReadTickets() ([]service.Ticket, error) {
	csvFile, err := os.Open(f.Path)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var Tickets []service.Ticket

	for _, record := range records {

		Ticket := service.Ticket{}
		Ticket.Id, err = strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}

		Ticket.Name = record[1]
		Ticket.Email = record[2]
		Ticket.Destination = record[3]
		Ticket.Date = record[4]
		Ticket.Price, err = strconv.Atoi(record[5])
		if err != nil {
			return nil, err
		}

		Tickets = append(Tickets, Ticket)

	}
	return Tickets, nil

}

func (f *File) Write(service.Ticket) error {
	return nil
}
