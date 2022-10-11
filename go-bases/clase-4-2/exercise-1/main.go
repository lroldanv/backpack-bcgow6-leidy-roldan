package main

import (
	"fmt"
	"os"
)

func main(){

	// defer func will be executed before execution is finished 
	defer func(){
		fmt.Println("ejecución finalizada")

	}()

	_, err := os.Open("customers.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
}