package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/", check)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":3060", nil))
}

func check(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("new-cookie")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "new-cookie",
			Value: "0",
			Path:  "/",
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	cookie.Value = strconv.Itoa(count)
	http.SetCookie(w, cookie)
	io.WriteString(w, cookie.Value)
}