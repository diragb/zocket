package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func (a *API) Add(w http.ResponseWriter, r *http.Request) {
	var song Song
	if err := decodeJSONBody(w, r, &song); err != nil {
		res, _ := json.Marshal(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}
	errResponse := Error{
		Status:    http.StatusBadRequest,
		Error:     http.StatusText(http.StatusBadRequest),
		Timestamp: time.Now(),
	}
	if song.Name == "" || song.Artists == nil || song.Album == "" {
		errResponse.Message = "Fields 'name', 'artists' and 'albums' cannot be empty"
		res, _ := json.Marshal(errResponse)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}
	song.Create()
	DB = append(DB, song)
	w.WriteHeader(http.StatusCreated)
}

func (a *API) GetAll(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal(DB)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (a *API) GetOne(w http.ResponseWriter, r *http.Request, s string) {
	for index := range DB {
		if DB[index].Slug == s {
			res, _ := json.Marshal(DB[index])
			w.WriteHeader(http.StatusOK)
			w.Write(res)
			return
		}
	}
	errResponse := Error{
		Status:    http.StatusNotFound,
		Error:     http.StatusText(http.StatusNotFound),
		Message:   "Song not found in database",
		Timestamp: time.Now(),
	}
	res, _ := json.Marshal(errResponse)
	w.WriteHeader(http.StatusBadRequest)
	w.Write(res)
}

func (a *API) Update(w http.ResponseWriter, r *http.Request, s string) {
	var updatedSong Song
	if err := decodeJSONBody(w, r, &updatedSong); err != nil {
		res, _ := json.Marshal(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}
	errResponse := Error{Timestamp: time.Now()}
	if updatedSong.Name == "" && updatedSong.Artists == nil && updatedSong.Album == "" {
		errResponse.Status = http.StatusBadRequest
		errResponse.Error = http.StatusText(http.StatusBadRequest)
		errResponse.Message = "Fields 'name', 'artists' and 'albums' cannot be empty"
		res, _ := json.Marshal(errResponse)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}
	// Iterate through the temporary database to find the song by slug
	// and update the song
	for index := range DB {
		if DB[index].Slug == s {
			DB[index].Update(updatedSong)
			res, _ := json.Marshal(DB[index])
			w.WriteHeader(http.StatusOK)
			w.Write(res)
			return
		}
	}
	errResponse.Status = http.StatusNotFound
	errResponse.Error = http.StatusText(http.StatusNotFound)
	errResponse.Message = "Song not found in database"
	res, _ := json.Marshal(errResponse)
	w.WriteHeader(http.StatusBadRequest)
	w.Write(res)
}

func (a *API) Delete(w http.ResponseWriter, r *http.Request, s string) {
	for index := range DB {
		if DB[index].Slug == s {
			DB = append(DB[:index], DB[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	errResponse := Error{
		Status:    http.StatusNotFound,
		Error:     http.StatusText(http.StatusNotFound),
		Message:   "Song not found in database",
		Timestamp: time.Now(),
	}
	res, _ := json.Marshal(errResponse)
	w.WriteHeader(http.StatusBadRequest)
	w.Write(res)
}
