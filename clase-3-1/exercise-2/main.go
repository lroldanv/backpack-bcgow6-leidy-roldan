package main

import (
	"encoding/csv"
	"fmt"
	"os"	
)

func main(){

	csvFile, err := os.Open("../myFile.csv")
	if err != nil {
		fmt.Println("Oops")
	}
	fmt.Println("Successfuly opened CSV file")

	r := csv.NewReader(csvFile)
	r.Comma = ';'

	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("Oops")
	}

	fmt.Printf("%s\t%10s\t%10s\n", "Id", "Price", "Quantity")
	for _, record := range records{
		fmt.Printf("%s\t%10s\t%10s\n", record[0], record[1], record[2])
	}
	
}