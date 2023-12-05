package models

import "time"

type Card struct {
	Number         string    `json:"number"`
	Cvv            string    `json:"cvv"`
	ExpirationDate time.Time `json:"expirationDate"`
	Balance        float64   `json:"balance"`
	CurrencyTag    string    `json:"CurrencyTag"`
	History        []int     `json:"history"`
	IsCardActive   bool      `json:"isCardActive"`
	Type           CardTypes `json:"cardType"`
}

type CardTypes string

const (
	CARD   CardTypes = "debit"
	CREDIT CardTypes = "credit"
)
