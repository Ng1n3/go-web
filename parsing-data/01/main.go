package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":3060", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	body := `
    	<form method="POST">
	 <input type="text" name="q">
	 <input type="submit">
	</form>
  `

	io.WriteString(w, body+v)
}
