package main

import (
	"log"
	"net/http"
)

func main() {
	// Creates an empty servemux
	mux := http.NewServeMux()

	// Creates a handler which 307 (temperorily redirect) redirects all requests it recieves to http://exmaple.org
	rh := http.RedirectHandler("http://example.org", 307)

	mux.Handle("/foo", rh)

	log.Println("Listening...")

	http.ListenAndServe(":3000", mux)

}
