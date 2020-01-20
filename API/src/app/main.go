package main

import (
	"app/responder"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", responder.SetHandler()))
}
