package main

import (
	"fmt"
	"net/http"
)

func GetSongHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here is the playlist!")
}

func PostSongHandler(w http.ResponseWriter, r *http.Request) {}

func PutSongHandler(w http.ResponseWriter, r *http.Request) {}

func DeleteSongHandler(w http.ResponseWriter, r *http.Request) {}
