// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type AttributeFilter struct {
	Colors       []int `json:"colors"`
	Origins      []int `json:"origins"`
	Sizes        []int `json:"sizes"`
	Prices       []int `json:"prices"`
	Availability []int `json:"availability"`
}

type AttributesFilter struct {
	Colors       []int `json:"colors"`
	Origins      []int `json:"origins"`
	Sizes        []int `json:"sizes"`
	Availability []int `json:"availability"`
	Prices       []int `json:"prices"`
}

type DetailItem struct {
	ID                int                 `json:"id"`
	Name              string              `json:"name"`
	Description       *string             `json:"description"`
	Order             int                 `json:"order"`
	Attributes        []*OverviewLabel    `json:"attributes"`
	Images            []*OverviewImage    `json:"images"`
	Collection        *OverviewCollection `json:"collection"`
	OrderInCollection int                 `json:"orderInCollection"`
}

type ItemAttributes struct {
	Colors       []*OverviewLabel `json:"colors"`
	Origins      []*OverviewLabel `json:"origins"`
	Sizes        []*OverviewLabel `json:"sizes"`
	Prices       []*OverviewLabel `json:"prices"`
	Availability []*OverviewLabel `json:"availability"`
}

type ListDetailItem struct {
	Data      []*DetailItem `json:"data"`
	Page      int           `json:"page"`
	Size      int           `json:"size"`
	Total     int           `json:"total"`
	IsEndPage bool          `json:"isEndPage"`
}

type ListItem struct {
	Data      []*OverviewItem  `json:"data"`
	Filter    *AttributeFilter `json:"filter"`
	Page      int              `json:"page"`
	Size      int              `json:"size"`
	Total     int              `json:"total"`
	IsEndPage bool             `json:"isEndPage"`
}

type NewCollection struct {
	Name string `json:"name"`
}

type NewImage struct {
	Link string `json:"link"`
}

type NewItem struct {
	ID          *int    `json:"id"`
	Name        *string `json:"name"`
	SearchKeys  *string `json:"searchKeys"`
	Description *string `json:"description"`
	Attributes  []int   `json:"attributes"`
	Images      []int   `json:"images"`
	Type        *int    `json:"type"`
}

type NewLabel struct {
	Code    string `json:"code"`
	Value   string `json:"value"`
	FkLabel *int   `json:"fkLabel"`
	Active  bool   `json:"active"`
}

type NewRole struct {
	Label       string  `json:"label"`
	Description *string `json:"description"`
	Active      bool    `json:"active"`
}

type NewUser struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Active      bool   `json:"active"`
	DisplayName string `json:"displayName"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	CountryCode int    `json:"countryCode"`
	Role        int    `json:"role"`
	Status      int    `json:"status"`
}

type OverviewCollection struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Order int    `json:"order"`
	Total int    `json:"total"`
}

type OverviewImage struct {
	ID   int    `json:"id"`
	Link string `json:"link"`
}

type OverviewItem struct {
	ID     int            `json:"id"`
	Name   string         `json:"name"`
	Avatar *OverviewImage `json:"avatar"`
	Price  *OverviewLabel `json:"price"`
}

type OverviewLabel struct {
	ID        int              `json:"id"`
	Code      string           `json:"code"`
	Value     string           `json:"value"`
	SubLabels []*OverviewLabel `json:"subLabels"`
}

type OverviewRole struct {
	ID          int     `json:"id"`
	Label       string  `json:"label"`
	Description *string `json:"description"`
}

type OverviewUser struct {
	ID          int            `json:"id"`
	Username    string         `json:"username"`
	Active      bool           `json:"active"`
	DisplayName string         `json:"displayName"`
	Address     string         `json:"address"`
	PhoneNumber string         `json:"phoneNumber"`
	CountryCode *OverviewLabel `json:"countryCode"`
	Role        *OverviewRole  `json:"role"`
	Status      *OverviewLabel `json:"status"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
}

type PaginationFilter struct {
	Collections []int             `json:"collections"`
	Page        int               `json:"page"`
	Size        int               `json:"size"`
	Keyword     *string           `json:"keyword"`
	Attributes  *AttributesFilter `json:"attributes"`
}
