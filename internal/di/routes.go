package di

import (
	"BankApi/internal/handlers"
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

func NewRouterContainer(service *ServiceContainer) *RouterContainer {
	return &RouterContainer{service: service}
}

func (c *RouterContainer) PostBill(ctx context.Context) *handlers.POSTBillsHandler {

	if c.postBill == nil {
		c.postBill = handlers.NewPOSTBillsHandler(c.service.CreateBill(ctx))

	}

	return c.postBill
}

func (c *RouterContainer) PostRegister(ctx context.Context) *handlers.POSTRegisterHandler {

	if c.postRegister == nil {
		c.postRegister = handlers.NewPOSTRegisterHandler(c.service.CreateUser(ctx))
	}
	return c.postRegister
}

func (c *RouterContainer) HTTPRouter(ctx context.Context) http.Handler {
	if c.router != nil {
		return c.router
	}

	router := mux.NewRouter()
	//router.Use(middleware.AuthMiddleware)

	router.Handle("/bills", c.PostBill(ctx)).Methods(http.MethodPost)
	router.Handle("/register", c.PostRegister(ctx)).Methods(http.MethodPost)

	c.router = router

	return c.router
}
