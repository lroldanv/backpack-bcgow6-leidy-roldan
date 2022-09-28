package main

import (
	"fmt"
	"math"
	// "sync"
)

type Product struct{
	Name string
	Price float64
	Quantity int
}

type Service struct{
	Name string
	Price float64
	TimeWorkedInMinutes float64
}

type Maintenance struct{
	Name string
	Price float64
}

func sumProducts(c chan float64, products...Product){
	var totalPrice float64

	for _, product := range products{

		totalPrice += product.Price * float64(product.Quantity)
	}
	c <- totalPrice // send sumProducts to c

}

func sumServices(c chan float64, services...Service){
	var totalPrice float64
	
	for _, service := range services{
		
		totalPrice +=  service.Price * math.Ceil(service.TimeWorkedInMinutes / 30)
	}
	c <- totalPrice // send sumServices to c
}

func sumMaintenances(c chan float64, maintenances...Maintenance){
	var totalPrice float64

	for _, maintenance := range maintenances{
		totalPrice += maintenance.Price
	}
	fmt.Println("sumMaintenances done")
	c <- totalPrice
}

func main() {


	p1 := Product{"cleaner", 2., 2}
	p2 := Product{"gasoline", 2., 2}

	s1 := Service{"clean", 2., 61.}
	s2 := Service{"dry", 2., 61.}

	m1 := Maintenance{"m1", 2.}
	m2 := Maintenance{"m2", 2.}

	c := make(chan float64)

	go sumProducts(c, p1, p2)
	go sumServices(c, s1, s2)
	go sumMaintenances(c, m1, m2)

	totalProducts, totalServices, totalMaintenances := <-c, <-c, <-c // receive from c
	 

	// var wg sync.WaitGroup
	// wg.Add(3)

	// go func(){
	// 	sumProducts(p1, p2)
	// 	wg.Done()
	// }()

	// go func(){
	// 	sumServices(s1,s2)
	// 	wg.Done()
	// }()

	// go func(){
	// 	sumMaintenances(m1,m2)
	// 	wg.Done()
	// }

	// wg.Wait()

	fmt.Println(totalProducts+totalServices+totalMaintenances)
	
}