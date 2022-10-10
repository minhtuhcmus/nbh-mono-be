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
