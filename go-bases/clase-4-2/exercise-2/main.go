package main

import(
	"fmt"
)

var FileID int = 1

type User struct{
	Name string
	LastName string
	DNI int
	Phone string
}

type File struct{
	User User
	FileID int
}

func (file File) assignFileID() File{
	file.FileID = FileID
	FileID++

	if FileID < 0{
		return nil
	}

	return file
}

func createUserRecord(user User) (userRecord string){
	userRecord = fmt.Sprintf("%s,%s,%d,%s", user.Name, user.LastName, user.DNI, user.Phone)
	return
}

func main(){
	file1 := File{}
	file2 := File{}

	if file.assignFileID()== nil{

	}
	
	fmt.Println(file1.assignFileID())
	fmt.Println(file2.assignFileID())
}



