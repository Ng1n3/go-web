package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func home(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Welcome home")
}

type dogList map[string][]string

func dogs(w http.ResponseWriter, req *http.Request) {
	dogs := dogList{
		"sheperd": []string{"german", "australian"},
		"bulldog": []string{"english", "french"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dogs)

}

func main() {
	http.HandleFunc("/", home)
	// http.Handle("/", http.HandlerFunc(home))
	http.HandleFunc("/dogs", (dogs))

	http.ListenAndServe(":3050", nil)
}
