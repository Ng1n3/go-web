package main

import (
	"log"
	"net/http"
)

func main() {
  
  // http.Handle("/", http.FileServer(http.Dir("./starting-files")))
  log.Fatal(http.ListenAndServe(":3060", http.FileServer(http.Dir("."))))
}