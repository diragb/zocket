package main

import "net/http"

// Songs CRUD API
// @Router	/add/				[POST]
// @Router	/get/				[GET]
// @Router	/get/{song-slug}	[GET]
// @Router	/update/{song-slug}	[PUT]
// @Router	/delete/{song-slug}	[DELETE]
func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &API{})
	if err := http.ListenAndServe(":8000", mux); err != nil {
		panic(err)
	}
}
