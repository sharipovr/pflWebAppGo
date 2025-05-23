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

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>Frequently Asked Questions</h1>
<ul>
  <li>
    <h3>Q: Is there a free version?</h3>
    <p>A: Yes! We offer a free trial for 30 days on any paid plans.</p>
  </li>

  <li>
    <h3>Q: What are your support hours?</h3>
    <p>A: We have support staff answering emails 24/7, though response times may be a bit slower on weekends.</p>
  </li>

  <li>
    <h3>Q: How do I contact support?</h3>
    <p>A: Email us – <a href="mailto:support@rustems.dev">support@rustems.dev</a></p>
  </li>
</ul>
	`)
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		http.Error(w, "Page not found", http.StatusNotFound)
// 	}
// }

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("Starting the server on: 3000...")
	http.ListenAndServe(":3000", router)
}
