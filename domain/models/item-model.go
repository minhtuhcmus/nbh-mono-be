package models

type Item struct {
	ID         int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string `json:"name"`
	SearchKeys string `json:"searchKeys"`
	Active     bool   `json:"active"`
}

type ItemImage struct {
	FkLabel int  `json:"fk_label"`
	FkItem  int  `json:"fk_item"`
	Order   int  `json:"order"`
	Active  bool `json:"active"`
}

type ItemAttribute struct {
	FkLabel int  `json:"fk_label"`
	FkItem  int  `json:"fk_item"`
	Active  bool `json:"active"`
}

type ItemAvatar struct {
	FkItem  int    `json:"fk_item"`
	FkImage int    `json:"fk_image"`
	Link    string `json:"link"`
}
