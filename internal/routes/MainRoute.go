package routes

import (
	"BankApi/internal/handlers"
	"fmt"
	"net/http"
)

type Route struct {
	accHandler *handlers.AccountHandler
}

func (route *Route) New(handler *handlers.AccountHandler) Route {
	fmt.Println(handler)
	return Route{accHandler: handler}
}

func (route *Route) Init(serverAddr string) {
	fmt.Println(route.accHandler)
	mux := http.NewServeMux()
	mux.HandleFunc("/register", route.accHandler.Registration)
	http.ListenAndServe(serverAddr, mux)

}
