package models

type User struct {
	ID int
	FirstName string
	LastName string
}

//variable block
var (
	//users - slice; *User - pointer
	//users is a slice that's going to hold pointers to User objects
	//collection of Users
	users []*User
	
	//At the package level there's no need to provide the colon operator to get the implicit initialization syntax.
	//Here, the compiler will imply the int data type.
	nextID = 1
)