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

//GetUsers return a slice of pointers to User objects
func GetUsers() []*User {
	return users;
}

//AddUser returns 2 values, the User created and a potential error object
func AddUser (u User) (User, error) {
	u.ID = nextID;
	nextID++;
	//use the address of operator (&) to grab a reference to the user object that came in. 
	//(The users collection is storing pointers to User objects, so we need the address of the user object that came in and store it)
	users = append(users, &u);
	return u, nil; //returning nil because no error occurrs here
}