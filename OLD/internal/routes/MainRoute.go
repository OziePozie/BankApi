package routes

import (
	handlers2 "BankApi/OLD/internal/handlers"
	"fmt"
	"net/http"
)

type Route struct {
	accHandler  *handlers2.AccountHandler
	billHandler *handlers2.BillHandler
	cardHandler *handlers2.CardHandler
}

func (route *Route) New() Route {
	handler := handlers2.New()
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
