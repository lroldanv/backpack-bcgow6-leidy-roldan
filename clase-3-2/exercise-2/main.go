package main

import "fmt"

type Product struct{
	Name string
	Price float64
	Quantity int
}

type User struct{
	Name string
	LastName string
	Email string
	Products []Product
}

// Create a new product
func addNewProduct(name string, price float64) Product{
	p := Product{}
	p.Name = name
	p.Price = price

	return p
}

// Add products to user list
func addProductsToUser(u *User, p *Product, quantity *int){
	p.Quantity = *quantity
	u.Products = append(u.Products, *p)	
}

// Delete products
func deleteUserProducts(u *User){
	u.Products = u.Products[:0]
}



func main(){
	fmt.Println("nyam")

	p1 := addNewProduct("taza", 2345.69)

	u := User{
		Name: "Pepito",
	}

	var q int

	addProductsToUser(&u, &p1, &q)

	fmt.Println(u)

	deleteUserProducts(&u)

	fmt.Println(u)

	

}