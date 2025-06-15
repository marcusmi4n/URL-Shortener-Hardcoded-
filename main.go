package main

import (
	"fmt"
	"net/http"
)

var links = map[string]string{
	"ggl": "https://google.com",
	"gh":  "https://github.com",
}

func handler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	url, ok := links[key]
	if ok {
		http.Redirect(w, r, url, http.StatusFound)
	} else {
		fmt.Fprintln(w, "Shortlink not found.")
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Shortener running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
