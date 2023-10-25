package routes

import (
	"BankApi/internal/handlers"
	"BankApi/internal/storage"
	"fmt"
	"net/http"
)

type Route struct {
	accHandler *handlers.AccountHandler
}

func (route *Route) New(storage *storage.Storage) Route {
	handler := handlers.New(storage)
	return Route{accHandler: handler}
}

func (route *Route) Init(serverAddr string) {
	fmt.Println(route.accHandler)
	mux := http.NewServeMux()
	mux.HandleFunc("/register", route.accHandler.Registration)
	mux.HandleFunc("/login", route.accHandler.Login)
	mux.HandleFunc("/accounts", route.accHandler.Accounts)

	http.ListenAndServe(serverAddr, mux)

}
