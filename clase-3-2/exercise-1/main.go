package main

import "fmt"

type user struct{
	name string
	lastName string
	age int
	email string
	password string
}

func changeNameFunc(u *user, newName string, newLastName string){
	u.name = newName
	u.lastName = newLastName
}

func (u *user) changeName(newName string, newLastName string){
	u.name = newName
	u.lastName = newLastName
}

func changeAgeFunc(u *user, newAge int){
	if newAge > 150 {
		fmt.Println("It's a vampire!")
	}
	u.age = newAge
}

func changeEmailFunc(u *user, newEmail string){
	u.email = newEmail
}

func changePasswordFunc(u *user, newPassword string){
	u.password = newPassword
}

func main() {
	u := user{
		name: "Valteri",
		lastName: "Bottas",
		age:32,
		email: "botitas@dev.com",
		password: "themostsecurepassword",
		
	}
	u.changeName("Jabber", "Wocky")

	fmt.Println("La direccion en memoria es: ", &u)

	changeNameFunc(&u, "Charles", "Hamilton")
	changeAgeFunc(&u, 250)
	changeEmailFunc(&u, "another@dev.co")
	changePasswordFunc(&u, "anoteherpassword")

	fmt.Println(u)
	fmt.Println("La direccion en memoria es: ", &u)

}