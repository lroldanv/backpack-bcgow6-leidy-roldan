package main

import (
	"fmt"
	"time"
)

// type admissionDateType struct {
// 	Year uint
// 	Month uint
// 	Day uint
// }

// String makes AdmissionDateType satisfy the stringer interface
// func(a admissionDateType) String() string{
// 	return fmt.Sprintf("%d/%d/%d", a.Year, a.Month, a.Day)
// }

type student struct {
	Name string
	LastName string
	Document uint
	AdmissionDate time.Time
}

func (s student) getDetails() {
	fmt.Println(s.Name)
	fmt.Printf("Name: %s \nLastName: %s \nDocument: %d \nAdmission Date: %s", s.Name, s.LastName, s.Document, s.AdmissionDate)
}

func main()  {

	s1 := student{
		Name: "Pepito",
		LastName: "Perez",
		Document: 1234,
		AdmissionDate: time.Date(2005, 2, 1, 12, 30, 0, 0, time.UTC),
	}
	
	s1.getDetails()
	
}