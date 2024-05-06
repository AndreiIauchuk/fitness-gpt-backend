package main

import (
	db "fitness-gpt-backend/internal/db"
	"log"
	"fmt"
	"net/http"
)

func main() {
	db.Init()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world")
	})

	err := http.ListenAndServe(":8080", nil)
	if (err != nil) {
		log.Fatal("Unable to start a web server!")
	}
}
