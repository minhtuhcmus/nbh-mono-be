package models

type Label struct {
	ID      int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Code    string `json:"code"`
	Value   string `json:"value"`
	FkLabel *int   `json:"fkLabel"`
	Active  bool   `json:"active"`
}
