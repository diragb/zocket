package main

import (
	"log"
	"net/http"
	"regexp"
	"strings"
)

var (
	create = regexp.MustCompile(`^\/add[\/]*$`)
	read   = regexp.MustCompile(`^(.*get[\\\/])`)
	update = regexp.MustCompile(`^(.*update[\\\/])`)
	delete = regexp.MustCompile(`^(.*delete[\\\/])`)
)

// Routes handler for the songs API
// Uses REGEX to match path and then executes the reqired functions
func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path)
	w.Header().Set("content-type", "application/json")
	path := strings.Split(r.URL.Path, "/")
	// Slug is used for representing song name
	slug := path[len(path)-1]
	switch {
	case create.MatchString(r.URL.Path):
		if r.Method == http.MethodPost {
			a.Add(w, r)
			return
		}
		methodNotAllowed(w, r)
	case read.MatchString(r.URL.Path):
		if r.Method == http.MethodGet {
			if slug != "" {
				a.GetOne(w, r, slug)
				return
			}
			a.GetAll(w, r)
			return
		}
		methodNotAllowed(w, r)
	case update.MatchString(r.URL.Path):
		if r.Method == http.MethodPut {
			if slug != "" {
				a.Update(w, r, slug)
				return
			}
			emptySlug(w, r)
			return
		}
		methodNotAllowed(w, r)
	case delete.MatchString(r.URL.Path):
		if r.Method == http.MethodDelete {
			if slug != "" {
				a.Delete(w, r, slug)
				return
			}
			emptySlug(w, r)
			return
		}
		methodNotAllowed(w, r)
	default:
		notFound(w, r)
		return
	}
}
