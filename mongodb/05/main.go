package main

import (
	"log"
	"net/http"
	"github.com/Ng1n3/go-web/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
  uc := controllers.NewUserController()
	r.GET("/", controllers.Index)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	log.Fatal(http.ListenAndServe(":3050", r))
}

