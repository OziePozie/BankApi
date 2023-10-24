package main

import (
	"BankApi/internal/storage"
	"log"
)

func main() {
	storage, err := storage.New()
	if err != nil {
		log.Fatal("Fail", err)
	}

}
