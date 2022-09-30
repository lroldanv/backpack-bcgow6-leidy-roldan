package service

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/google/uuid"
)

type Bookings interface {
	// Create create a new Ticket
	Create(name, email, destination, date string, price int) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
}

var CreateTicketError = errors.New("the ticket was not created")

type bookings struct {
	Tickets []Ticket
}

type Ticket struct {
	Id, Name, Email, Destination, Date string
	Price                              int
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func (b *bookings) Create(name, email, destination, date string, price int) (Ticket, error) {

	if len(name) < 1 {
		return Ticket{}, fmt.Errorf("error: a name must be provided, %w", CreateTicketError)
	}

	regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !(regex.MatchString(email)) {
		return Ticket{}, fmt.Errorf("error: invalid email, %w", CreateTicketError)
	}

	if len(destination) < 1 {
		return Ticket{}, fmt.Errorf("error: a destination must be provided, %w", CreateTicketError)
	}

	if len(destination) < 1 {
		return Ticket{}, fmt.Errorf("error: a destination must be provided, %w", CreateTicketError)
	}

	// TODO validate regex for date format
	if len(date) < 1 {
		return Ticket{}, fmt.Errorf("error: a date must be provided, %w", CreateTicketError)
	}

	NewTicket := Ticket{}
	NewTicket.Id = uuid.NewString()
	NewTicket.Name = name
	NewTicket.Email = email
	NewTicket.Destination = destination
	NewTicket.Date = date
	NewTicket.Price = price

	return NewTicket, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	return Ticket{}, nil
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	return Ticket{}, nil
}

func (b *bookings) Delete(id int) (int, error) {
	return 0, nil
}
