package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type Person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {

	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":3060", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	//body
	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)
	fmt.Println(bs)
	body := string(bs)

	err := tpl.ExecuteTemplate(w, "index.gohtml", body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
