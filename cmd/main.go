package main

import (
	"BankApi/internal/routes"
)

func main() {

	mainRoute := routes.Route{}

	mainRoute = mainRoute.New()
	mainRoute.Init(":4000")

}
