package controllers

import (
	"regexp"
	"net/http"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

//variable name - uc
//type to bind to - userController
//method name - ServeHTTP. Requires a ResponseWriter object from the HTTP package and it's going to receive a Request object
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//convert string to byte slice
	w.Write([]byte("Hello from User Controller"));
}

//Go doesn't use constructors, it uses constructor functions
//Convention in naming a constructor function is starting with the word 'new', then the name of the type of object that will be constructed
//often returns a pointer to the object to avoid an unnecessary copy operation
func newUserController() *userController {
	//When dealing with struct data types, constructing an object and using the 'address of' operator (&) without a named variable is permissible. You can immediately take the address of it
	return &userController{
		//regex will look for paths that are '/users/ a number'
		userIDPattern: regexp.MustCompile(`/users/(\d+)/?`),
	}
	//This is actually a local variable. We're constructing userController in the scope of this function & returning the address of it
	//Go will recognize if we're returning the address of a local variable & automatically promote the variable up to the level that it needs to be so we don't lose the variable & we don't lose track of the memory that was being allocated by that function
}