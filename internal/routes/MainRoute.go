package routes

import (
	"BankApi/internal/handlers"
	"fmt"
	"net/http"
)

type Route struct {
	accHandler *handlers.AccountHandler
}

func (route *Route) New() Route {

	return Route{accHandler: handlers.New()}
}

func (route *Route) Init(serverAddr string) {
	fmt.Println(route.accHandler)
	mux := http.NewServeMux()
	mux.HandleFunc("/register", route.accHandler.Registration)
	mux.HandleFunc("/login", route.accHandler.Login)
	http.ListenAndServe(serverAddr, mux)

}
