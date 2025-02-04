package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func main() {
  http.HandleFunc("/", index)
  http.Handle("/favicon.ico", http.NotFoundHandler())
  log.Fatal(http.ListenAndServe(":3060", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
  cookie, err := req.Cookie("session")
  if err != nil {
    id, err := uuid.NewV7()
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    cookie = &http.Cookie{
      Name: "session",
      Value: id.String(),
      HttpOnly: true,
      Path: "/",
    }
    http.SetCookie(w, cookie)
  }
  fmt.Println(cookie)
}