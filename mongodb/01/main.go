package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)


func main() {
  r := httprouter.New()
  r.GET("/", index)
  log.Fatal(http.ListenAndServe("3050", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
  fmt.Fprint(w, "Welcome!\n")
}