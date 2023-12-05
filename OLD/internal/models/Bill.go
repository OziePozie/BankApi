package models

type Bill struct {
	ID           int    `json:"ID"`
	Number       string `json:"number"`
	Limit        int    `json:"limit"`
	Cards        []int  `json:"cards_id"`
	History      []int  `json:"history_id"`
	IsBillActive bool   `json:"isBillActive"`
}
