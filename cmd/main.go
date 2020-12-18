package main

import (
	"log"
	"net/http"

	"github.com/tatsuya4559/bookreviews"
)

func main() {
	bookreviews.BuildContainer()

	router := bookreviews.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
