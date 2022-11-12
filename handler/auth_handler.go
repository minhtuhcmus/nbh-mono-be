package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/minhtuhcmus/nbh-mono-be/domain/models"
	"github.com/minhtuhcmus/nbh-mono-be/domain/services"
	"github.com/minhtuhcmus/nbh-mono-be/utils"
	"net/http"
	"time"
)

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

func NewAuthHandler(authService *services.AuthService) func(r chi.Router) {

	ah := &authHandler{authService: *authService}

	return func(r chi.Router) {
		r.Post("/signin", ah.SignIn)
		r.Post("/signup", ah.SignUp)
		r.Post("/verify-token", ah.VerifyToken)
	}
}

type authHandler struct {
	authService services.AuthService
}

type AuthHandler interface {
	SignIn(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
	VerifyToken(w http.ResponseWriter, r *http.Request)
}

func (ah *authHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	t := time.Now().UnixNano()

	var creds Credentials
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	accessToken, err := ah.authService.SignIn(r.Context(), creds.Username, creds.Password)
	if err != nil {
		http.Error(w, fmt.Sprintf("Cannot sign in %v", err), http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Domain: "http://localhost:8080",
		Path:   "/",
		Name:   "access_token",
		Value:  accessToken,
		Secure: true,
		MaxAge: 300,
	})

	fmt.Println("SignIn Handler ", time.Now().UnixNano()-t)

}

func (ah *authHandler) SignUp(w http.ResponseWriter, r *http.Request) {
}

func (ah *authHandler) VerifyToken(w http.ResponseWriter, r *http.Request) {
	authCookie, err := r.Cookie("access_token")
	if err != nil {
		json.NewEncoder(w).Encode(models.TokenAuthenticationResponse{
			IsAuth:  false,
			IsAdmin: false,
		})
		return
	}
	if authCookie.Value == "" {
		json.NewEncoder(w).Encode(models.TokenAuthenticationResponse{
			IsAuth:  false,
			IsAdmin: false,
		})
		return
	}

	authClaims, err := utils.JwtValidate(authCookie.Value)
	if err != nil {
		json.NewEncoder(w).Encode(models.TokenAuthenticationResponse{
			IsAuth:  false,
			IsAdmin: false,
		})
		return
	}

	if authClaims == nil || authClaims.Roles == nil {
		json.NewEncoder(w).Encode(models.TokenAuthenticationResponse{
			IsAuth:  false,
			IsAdmin: false,
		})
		return
	}

	json.NewEncoder(w).Encode(models.TokenAuthenticationResponse{
		IsAuth:  true,
		IsAdmin: utils.IsAdmin(authClaims.Roles),
	})
}
