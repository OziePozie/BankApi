package routes

import (
	"BankApi/internal/handlers"
	"fmt"
	"net/http"
)

type Route struct {
	accHandler  *handlers.AccountHandler
	billHandler *handlers.BillHandler
	cardHandler *handlers.CardHandler
}

func (route *Route) New() Route {
	handler := handlers.New()
	return Route{
		accHandler:  handler.AccHandler,
		billHandler: handler.BillHandler,
		cardHandler: handler.CardHandler,
	}
}

func (route *Route) Init(serverAddr string) {
	fmt.Println(route.accHandler)
	mux := http.NewServeMux()
	mux.HandleFunc("/login", route.accHandler.Login)
	mux.HandleFunc("/accounts", route.accHandler.Accounts)
	mux.HandleFunc("/bills", route.billHandler.Bills)
	//mux.HandleFunc("/bills", route.billHandler.CreateBill)
	//mux.HandleFunc("/bills/set-limit", route.billHandler.SetLimit)
	mux.HandleFunc("/cards", route.cardHandler.CreateCard)
	http.ListenAndServe(serverAddr, mux)

}
