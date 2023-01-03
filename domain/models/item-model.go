package models

type Item struct {
	ID         int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string `json:"name"`
	SearchKeys string `json:"searchKeys"`
	Active     bool   `json:"active"`
}

type ItemImage struct {
	FkItem   int  `json:"fk_item"`
	FkImage  int  `json:"fk_image"`
	Order    int  `json:"order"`
	IsAvatar bool `json:"is_avatar"`
	Active   bool `json:"active"`
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

type ItemCollection struct {
	FkItem       int  `json:"fkItem"`
	FkCollection int  `json:"fkCollection"`
	Active       bool `json:"active"`
}

type DetailItem struct {
	ID                int     `json:"id"`
	Name              string  `json:"name"`
	Description       *string `json:"description"`
	Order             int     `json:"order"`
	Attributes        *string `json:"attributes"`
	Images            *string `json:"images"`
	Collection        *string `json:"collection"`
	OrderInCollection int     `json:"orderInCollection"`
}

type OverviewItem struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Avatar *string `json:"avatar"`
	Price  *string `json:"price"`
}

type ItemAttributeWithSubLabels struct {
	ID     int    `json:"id"`
	Code   string `json:"code"`
	Value  string `json:"value"`
	Labels string `json:"labels"`
}

type AttributeFilter struct {
	Colors       string `json:"colors"`
	Origins      string `json:"origins"`
	Sizes        string `json:"sizes"`
	Availability string `json:"availability"`
	Prices       string `json:"prices"`
}
