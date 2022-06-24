package main

import (
	"net/http"

	"github.com/JoseLuizFranco/lyricsapi/api"
)

func main() {
	http.HandleFunc("/api/lyrics", handler.Lyrics)
	http.ListenAndServe(":8080", nil)
}
