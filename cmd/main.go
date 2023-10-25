package main

import (
	"BankApi/internal/handlers"
	"BankApi/internal/repository"
	"BankApi/internal/routes"
	"BankApi/internal/storage"
	"fmt"
)

func main() {

	mainRoute := routes.Route{}
	s, _ := storage.New()
	k := repository.New(s)

	h := handlers.New(k)
	fmt.Println(h)
	mainRoute = mainRoute.New(h)
	mainRoute.Init(":4000")

}
