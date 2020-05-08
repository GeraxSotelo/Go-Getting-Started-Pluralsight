package controllers

import (
	"net/http"
)

func RegisterControllers() {
	uc := newUserController()

	//Any object, any type, that has a method on it called ServeHTTP & has the correct function signature will implement this Handler interface
	http.Handle("/users", *uc);
	//When working with routing, "/users" is different than "/users/"
	http.Handle("/users/", *uc);
	//"/users/" will handle "/users/plus anything else"
}