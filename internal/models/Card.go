package models

import "time"

type Card struct {
	Number         string    `json:"number"`
	Cvv            string    `json:"cvv"`
	ExpirationDate time.Time `json:"expirationDate"`
	Balance        float64   `json:"balance"`
	CurrencyTag    string    `json:"CurrencyTag"`
	History        []History `json:"history"`
	IsCardActive   bool      `json:"isCardActive"`
}
