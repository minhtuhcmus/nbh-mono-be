package models

type User struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	HashedPassword string `json:"hashedPassword"`
	Active         bool   `json:"active"`
}

type TokenAuthenticationResponse struct {
	IsAuth  bool `json:"isAuth"`
	IsAdmin bool `json:"isAdmin"`
}

type UserRole string

const (
	Admin         UserRole = "admin"
	SaleManager   UserRole = "sale-manager"
	SalePerson    UserRole = "sale-person"
	Packer        UserRole = "packer"
	PackerManager UserRole = "packer-manager"
	Accountant    UserRole = "accountant"
	Shipper       UserRole = "shipper"
	Customer      UserRole = "customer"
)

type UserAccessControlList struct {
	ID    int
	Roles []UserRole
}
