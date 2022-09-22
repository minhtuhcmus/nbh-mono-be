package models

type Collection struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Order  int    `json:"order"`
	Active bool   `json:"active"`
}
