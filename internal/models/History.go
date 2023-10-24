package models

import "time"

type History struct {
	ID                int       `json:"ID"`
	Date              time.Time `json:"date"`
	DestinationCardId string    `json:"destination_Card_ID"`
	ArrivalCardId     string    `json:"arrival_Card_ID"`
	OperationType     string    `json:"operationType"`
	Sum               float64   `json:"sum"`
}
