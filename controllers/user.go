package controllers

import (
	"net/http"
)

type userController struct {}

//variable name - uc
//type to bind to - userController
//method name - ServeHTTP. Requires a ResponseWriter object from the HTTP package and it's going to receive a Request object
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//convert string to byte slice
	w.Write([]byte("Hello from User Controller"));
}