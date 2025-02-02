package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":3060", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	var s string

	fmt.Println(req.Method)
	if req.Method == http.MethodPost {
		f, h, err := req.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		defer f.Close()

		fmt.Println("\nfile:", f, "\nheader:", h, "\nerror:", err)

		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}
	w.Header().Set("Content-Type", "text/html charset=utf-8")

	body := `
    	<form method="POST" enctype="multipart/form-data">
	 <input type="text" name="q">
	 <input type="submit">
	</form>
  `

	io.WriteString(w, body+s)
}
