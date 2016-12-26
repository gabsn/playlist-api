package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/songs", GetSongHandler).Methods("GET")
	r.HandleFunc("/api/songs", PostSongHandler).Methods("POST")
	r.HandleFunc("/api/songs/{id}", PutSongHandler).Methods("PUT")
	r.HandleFunc("/api/songs/{id}", DeleteSongHandler).Methods("DELETE")

	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		WriteTimeout:   10 * time.Second,
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Listening...")
	log.Fatal(server.ListenAndServe())
}
