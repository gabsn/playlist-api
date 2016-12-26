package main

import (
	"encoding/json"
	"net/http"
	"time"
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
	w.Write(j)
}

// HTTP POST /api/songs
func PostSongHandler(w http.ResponseWriter, r *http.Request) {
	var song Song
	err := json.NewDecoder(r.Body).Decode(&song)
	if err != nil {
		panic(err)
	}
	song.AddedOn = time.Now()
	k := song.AddedOn.String()
	playlist[k] = song
}

// HTTP PUT /api/songs/{id}
func PutSongHandler(w http.ResponseWriter, r *http.Request) {}

// HTTP DELETE /api/songs/{id}
func DeleteSongHandler(w http.ResponseWriter, r *http.Request) {}
