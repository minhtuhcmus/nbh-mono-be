package models

import "time"

type Stock struct {
	FkItem        int       `json:"fk_item"`
	Quantity      int       `json:"quantity"`
	AvailableFrom time.Time `json:"available_from"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Active        bool      `json:"active"`
}

type StockLogAction string

const (
	Add  StockLogAction = "add"
	Subs StockLogAction = "subs"
)

type StockLog struct {
	FkStock      int            `json:"fk_stock"`
	ChangeAmount int            `json:"change_amount"`
	Action       StockLogAction `json:"action"`
	Note         string         `json:"note"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	Active       bool           `json:"active"`
}

type StockAmount struct {
	FkItem         int `json:"fk_item"`
	AvailableStock int `json:"available_stock"`
}
