package main

import (
	"net/http"
	"os"
)

// Songs CRUD API
// @Router	/add/				[POST]
// @Router	/get/				[GET]
// @Router	/get/{song-slug}	[GET]
// @Router	/update/{song-slug}	[PUT]
// @Router	/delete/{song-slug}	[DELETE]
func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &API{})
	port := os.Getenv("PORT")
	if err := http.ListenAndServe("0.0.0.0:"+port, mux); err != nil {
		panic(err)
	}
}
