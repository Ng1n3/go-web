package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ng1n3/go-web/models"
	"github.com/julienschmidt/httprouter"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func Index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
  <html lang="en">
  <head>
  <meta charset="UTF-8">
  <title>Index</title>
  </head>
  <body>
  <a href="/user/9872309847">GO TO: http://localhost:3050/user/9872309847</a>
  </body>
  </html>
  `
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

func GetUser(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "Tom",
		Gender: "male",
		Age:    9,
		Id:     p.ByName("id"),
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func CreateUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(req.Body).Decode(&u)

	u.Id = "007"

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func DeleteUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	//TODO: write code to delete user
	var users = map[string]models.User{
		"9872309847": {Name: "Tom", Gender: "male", Age: 9, Id: "9872309847"},
		"1234567890": {Name: "Alice", Gender: "female", Age: 25, Id: "1234567890"},
	}

	id := p.ByName("id")
	_, ok := users[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User with ID %s not found\n", id)
		return
	}

	delete(users, id)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User with ID %s deleted successfully\n", id)
}
