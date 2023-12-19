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
	postLogin    *handlers.POSTLoginHandler
	getBill      *handlers.GETBillsHandler
	postDeposit  *handlers.POSTDepositHandler
}

func (c *RouterContainer) PostDeposit(ctx context.Context) *handlers.POSTDepositHandler {

	if c.postDeposit == nil {
		c.postDeposit = handlers.NewPOSTDepositHandler(c.service.Deposit(ctx))
	}

	return c.postDeposit
}

func (c *RouterContainer) PostLogin(ctx context.Context) *handlers.POSTLoginHandler {

	if c.postLogin == nil {
		c.postLogin = handlers.NewPOSTLoginHandler(c.service.LoginUser(ctx))
	}

	return c.postLogin
}

func (c *RouterContainer) SetPostLogin(postLogin *handlers.POSTLoginHandler) {
	c.postLogin = postLogin
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

	router.Use(middleware.PanicRecovery)

	billRouter := router.PathPrefix("/api").Subrouter()
	billRouter.Use(middleware.AuthMiddleware)

	//router.Handle("/bills", c.PostBill(ctx)).Methods(http.MethodPost)
	billRouter.Handle("/bills", c.PostBill(ctx)).Methods(http.MethodPost)
	billRouter.Handle("/bills", c.GetBill(ctx)).Methods(http.MethodGet)
	billRouter.Handle("/deposits", c.PostDeposit(ctx)).Methods(http.MethodPost)

	router.Handle("/register", c.PostRegister(ctx)).Methods(http.MethodPost)
	router.Handle("/login", c.PostLogin(ctx)).Methods(http.MethodPost)
	c.router = router

	return c.router
}

func (c *RouterContainer) GetBill(ctx context.Context) http.Handler {
	if c.getBill == nil {
		c.getBill = handlers.NewGETBillsHandler(c.service.GetBill(ctx))

	}

	return c.getBill

}
