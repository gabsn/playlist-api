package main

import (
	"encoding/json"
	"net/http"
)

// HTTP GET /api/songs
func GetSongHandler(w http.ResponseWriter, r *http.Request) {
	var songs []Song
	for _, v := range playlist {
		songs = append(songs, v)
    }
	w.Header().Set("Content-Type", "apllication/json")
    j, err := json.Marshal(songs)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// HTTP POST /api/songs
func PostSongHandler(w http.ResponseWriter, r *http.Request) {}

// HTTP PUT /api/songs/{id}
func PutSongHandler(w http.ResponseWriter, r *http.Request) {}

// HTTP DELETE /api/songs/{id}
func DeleteSongHandler(w http.ResponseWriter, r *http.Request) {}
