package main

import (
	di "BankApi/internal/di"
	"context"
	"log"
	"net/http"
)

func main() {
	container := di.NewContainer()
	ctx := context.Background()

	err := http.ListenAndServe(":8080", container.Routes().HTTPRouter(ctx))
	if err != nil {
		log.Fatal(err)
	}
}
