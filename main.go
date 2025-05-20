package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my web site built with Go !</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact page</h1><p>To get in touch email me at <a href=\"mailto:sharipov.rustem@gmail.com\">sharipov.rustem@gmail.com</a></p>")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		// TO DO:
		// handle the page not found error
	}
}

func main() {
	// http.HandleFunc("/", homeHandler)
	http.HandleFunc("/", pathHandler)
	// http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting the server on: 3000...")
	http.ListenAndServe(":3000", nil)
}
