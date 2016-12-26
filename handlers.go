package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func WriteToJSON(w http.ResponseWriter, object interface{}, status int) {
    w.Header().Set("Content-Type", "apllication/json")
	j, err := json.Marshal(object)
	if err != nil {
		panic(err)
	}
    w.WriteHeader(status)
	w.Write(j)
}

// HTTP GET /api/songs
func GetSongHandler(w http.ResponseWriter, r *http.Request) {
	var songs []Song
	for _, v := range playlist {
		songs = append(songs, v)
	}
    WriteToJSON(w, songs, http.StatusOK)
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
    WriteToJSON(w, song, http.StatusCreated)
}

// HTTP PUT /api/songs/{id}
func PutSongHandler(w http.ResponseWriter, r *http.Request) {}

// HTTP DELETE /api/songs/{id}
func DeleteSongHandler(w http.ResponseWriter, r *http.Request) {}
