package main

import (
	"log"
	"net/http"
	"github.com/Ng1n3/go-web/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", controllers.Index)
	r.GET("/user/:id", controllers.GetUser)
	r.POST("/user", controllers.CreateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)
	log.Fatal(http.ListenAndServe(":3050", r))
}

