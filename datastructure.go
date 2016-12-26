package main

import (
	"time"
)

type Song struct {
	Title   string    `json:"title"`
	Artist  string    `json:"artist"`
	AddedOn time.Time `json:"addedon"`
}

var playlist = make(map[string]Song)

var id int = 0
