package main

import "fmt"

func main(){
	word := "esparrago"
	fmt.Println(len(word))

	for i := 0; i < len(word); i++{
		fmt.Println(string(word[i]))
	}

	
}