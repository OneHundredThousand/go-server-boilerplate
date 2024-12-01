package main

import (
	"log"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
