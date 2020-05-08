package models

import (
	"fmt"
	"errors"
)

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
	if u.ID != 0 {
		//Expecting a User value to be returned. We can return an empty User object. Can't return nil because we're not returning a pointer value 
		return User{}, errors.New("New User must not include id or it must be set to zero");
	}
	u.ID = nextID;
	nextID++;
	//use the address of operator (&) to grab a reference to the user object that came in. 
	//(The users collection is storing pointers to User objects, so we need the address of the user object that came in and store it)
	users = append(users, &u);
	return u, nil; //returning nil because no error occurrs here
}

func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			//return user, dereference the pointer because we're expecting a user value to come back out
			return *u, nil; //nil is a signal to the caller that it got a valid user since no error was provided
		}
	}
	//fmt.Errorf() allows you to return formatted strings as error objects
	return User{}, fmt.Errorf("User with ID '%v' not found", id);
}

func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u;
			return u, nil;
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID);
}

func RemoveUserById(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...);
			return nil;
		}
	}
	return fmt.Errorf("User with ID '%v' not found", id);
}