package controllers

import (
	"strconv"
	"github.com/pluralsight/go-gettingstarted/webservice/models"
	"encoding/json"
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
//Because our built controller has a method ServeHTTP (same name as handler interface), & receives a ResponseWriter & a pointer to a Request object as parameters, it automatically implements the handler interface. We don't need to specify that we're explicitly implementing that interface

	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w, r);
		case http.MethodPost:
			uc.post(w, r);
		default:
			w.WriteHeader(http.StatusNotImplemented);
		}
	} else {
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		//convert string to numerical data type
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound);
		}
		switch r.Method {
		case http.MethodGet:
			uc.get(id, w)
		case http.MethodPut:
			uc.put(id, w, r)
		case http.MethodDelete:
			uc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}

	//convert string to byte slice
	w.Write([]byte("Hello from User Controller"));
}

func (uc *userController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetUsers(), w);
}

func (uc *userController) get(id int, w http.ResponseWriter) {
	u, err := models.GetUserByID(id);
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError);
		return;
	}
	encodeResponseAsJSON(u, w); //take User object & convert to a JSON representation
}

func (uc *userController) post(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r);
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError);
		w.Write([]byte("Could not parse User object"));
		return;
	}
	u, err = models.AddUser(u);
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError);
		w.Write([]byte(err.Error()));
		return;
	}
	encodeResponseAsJSON(u, w);
}

func (uc *userController) put(id int, w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r);
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError);
		w.Write([]byte("Could not parse User object"));
		return;
	}
	if id != u.ID {
		w.WriteHeader(http.StatusBadRequest);
		w.Write([]byte("ID of submnitted user must match ID in URL"));
		return;
	}
	u, err = models.UpdateUser(u);
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError);
		w.Write([]byte(err.Error()));
		return;
	}
	encodeResponseAsJSON(u, w);
}

func (uc *userController) delete(id int, w http.ResponseWriter) {
	err := models.RemoveUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError);
		w.Write([]byte(err.Error()));
		return;
	}
	w.WriteHeader(http.StatusOK);
}

//take JSON body and convert it into a User object
func (uc *userController) parseRequest(r *http.Request) (models.User, error) {
	dec := json.NewDecoder(r.Body);
	var u models.User;
	err := dec.Decode(&u);
	if err != nil {
		return models.User{}, err;
	}
	return u, nil;
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