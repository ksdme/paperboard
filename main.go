package main

import (
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	io.WriteString(w, "hello world")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)

	bind := ":8080"
	log.Println("Starting paperboard on", bind)
	log.Fatal(http.ListenAndServe(bind, router))
}
