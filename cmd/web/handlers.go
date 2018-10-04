package main

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("./ui/html/index.html"))

// The different response values that a site check can result in.
const (
	badRequest      byte = 1
	siteUnreachable byte = 2
	siteReachable   byte = 3
)

// Specifically handles responses for the index page ("/").
type indexPage struct {
	Status byte
	Domain string
}

// IndexRouter routes the incoming request for the index page ("/")
// and performs the appropriate action depending on the Method used.
func IndexRouter(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method == "GET" {
		defaultIndexHandler(w, r)
		return
	}
	if r.Method == "POST" {
		checkHandler(w, r)
		return
	}

	http.NotFound(w, r)
}

// defaultIndexHandler renders the default index page layout.
func defaultIndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}

// checkHandler performs a GET request for the domain name provided by the sender
// and returns a response indicating if it was successful or not.
func checkHandler(w http.ResponseWriter, r *http.Request) {
	d := r.FormValue("domain")

	if !IsValidDomain(d) {
		tmpl.Execute(w, indexPage{
			badRequest,
			d,
		})
		return
	}

	if !IsSiteUp(d) {
		tmpl.Execute(w, indexPage{
			siteUnreachable,
			d,
		})
		return
	}

	tmpl.Execute(w, indexPage{
		siteReachable,
		d,
	})
}
