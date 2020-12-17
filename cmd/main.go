package main

import (
	"log"
	"net/http"

	"github.com/tatsuya4559/bookreviews"
	"github.com/tatsuya4559/bookreviews/models"
)

func main() {
	err := models.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	router := bookreviews.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
