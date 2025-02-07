package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Ng1n3/go-web/models"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct{
	Client *mongo.Client
}

func NewUserController(s *mongo.Client) *UserController {
	return &UserController{s}
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

func (uc UserController) GetUser(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "Tom",
		Gender: "male",
		Age:    9,
		Id:     primitive.NewObjectID(),
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	u := models.User{}
	if err := json.NewDecoder(req.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u.Id = primitive.NewObjectID()
	collection := uc.Client.Database("persona").Collection("users")
	_, err := collection.InsertOne(ctx, u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	uj, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController)DeleteUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()


	id := p.ByName("id")
	objID, err := primitive.ObjectIDFromHex("id")
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	collection := uc.Client.Database("persona").Collection("users")
	result, err := collection.DeleteOne(ctx, primitive.M{"_id": objID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if result.DeletedCount == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User with ID %s deleted successfully\n", id)
}
