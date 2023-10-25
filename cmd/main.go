package main

import (
	"BankApi/internal/routes"
	storage "BankApi/internal/storage"
)

func main() {

	mainRoute := routes.Route{}
	storage, _ := storage.New()
	mainRoute = mainRoute.New(storage)
	mainRoute.Init(":4000")

}
