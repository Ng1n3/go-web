package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type user struct {
  UserName string
  First string
  Last string
}

var tpl *template.Template
var dbUsers = map[string]user{} //user Id, user
var dbSessions = map[string]string{} // session ID, userId

func init() {
  tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
  http.HandleFunc("/", index)
  http.HandleFunc("/bar", bar)
  http.Handle("/favicon.ico", http.NotFoundHandler())
  log.Fatal(http.ListenAndServe(":3050", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
  c, err := req.Cookie("session")
  if err != nil {
    sID, err := uuid.NewV7()
    if err != nil {
      http.Error(w, "error with uuid", http.StatusInternalServerError)
    }

    c = &http.Cookie{
      Name: "session",
      Value: sID.String(),
    }
    http.SetCookie(w, c)
  }

  var u user
  if un, ok := dbSessions[c.Value]; ok {
    u = dbUsers[un]
  }

  if req.Method == http.MethodPost{
    un := req.FormValue("username")
    f := req.FormValue("firstname")
    l := req.FormValue("lastname")
    u = user{un, f, l}
    dbSessions[c.Value] = un
    dbUsers[un] = u
  }

  tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
  c, err := req.Cookie("session")
  if err != nil {
    http.Redirect(w, req, "/", http.StatusSeeOther)
    return
  }

  un, ok := dbSessions[c.Value]
  if !ok {
    http.Redirect(w, req, "/",  http.StatusSeeOther)
    return
  }

  u := dbUsers[un]
  tpl.ExecuteTemplate(w, "bar.gohtml", u)
}