package main

import (
	"fmt"
	"log"
	"net/http"

	"./github.com/julienschmidt/httprouter"
	"./wordmixer"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Puzzle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	puzzle := wordmixer.NewPuzzle(ps.ByName("name"), 2)
	fmt.Fprint(w, wordmixer.ToHtml(puzzle))
}

func main() {
	log.Print("Starting WordMixer")
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/puzzle/:name", Puzzle)

	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
