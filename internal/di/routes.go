package di

import (
	"BankApi/internal/handlers"
	"BankApi/internal/handlers/middleware"
	"context"
	"github.com/gorilla/mux"
	"net/http"
)

type RouterContainer struct {
	service *ServiceContainer

	router http.Handler

	postBill     *handlers.POSTBillsHandler
	postRegister *handlers.POSTRegisterHandler
}

func (c *RouterContainer) HTTPRouter(ctx context.Context) http.Handler {
	if c.router != nil {
		return c.router
	}

	router := mux.NewRouter()
	router.Use(
		middleware.AuthMiddleware)

	router.Handle("/bills", c.postBill).Methods(http.MethodPost)
	router.Handle("/register", c.postRegister).Methods(http.MethodPost)

	c.router = router

	return c.router
}
