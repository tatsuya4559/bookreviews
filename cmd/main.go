package main

import (
	"log"
	"net/http"

	"github.com/tatsuya4559/bookreviews"
	"github.com/urfave/negroni"
)

func main() {
	// DI
	bookreviews.BuildContainer()

	// router
	router := bookreviews.NewRouter()

	// middleware
	n := negroni.Classic()
	n.UseHandler(router)

	log.Fatal(http.ListenAndServe(":8000", n))
}
