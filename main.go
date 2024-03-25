package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ksdme/paperboard/pages"
)

func main() {
	router := httprouter.New()
	router.GET("/", pages.Dashboard)

	bind := ":8080"
	log.Println("Starting paperboard on", bind)
	log.Fatal(http.ListenAndServe(bind, router))
}
