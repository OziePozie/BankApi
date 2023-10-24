package main

import (
	"BankApi/internal/server"
	"BankApi/internal/storage"
	"log"
	"net/http"
)

func main() {
	_, err := storage.New()
	if err != nil {
		log.Fatal("Fail", err)
	}

	s := &server.Server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
