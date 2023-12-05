package handlers

import "net/http"

type (
	RegisterRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	Register struct {
	}
)

func (Register) handle(w http.ResponseWriter, r *http.Request) {

}
