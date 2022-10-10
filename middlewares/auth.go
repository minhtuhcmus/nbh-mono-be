package middlewares

import (
	"context"
	"fmt"
	"github.com/minhtuhcmus/nbh-mono-be/database/datastore"
	"github.com/minhtuhcmus/nbh-mono-be/utils"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

func (m middleware) WithAuth() func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")

			if auth == "" {
				//http.Error(w, fmt.Sprintln("Authorization header not found"), http.StatusUnauthorized)
				next.ServeHTTP(w, r)
				return
			}

			bearer := "Bearer "
			auth = auth[len(bearer):]

			authClaims, err := utils.JwtValidate(auth)
			if err != nil {
				http.Error(w, fmt.Sprintf("Invalid token %v", err), http.StatusForbidden)
				return
			}

			if authClaims.StandardClaims.ExpiresAt < time.Now().Unix() {
				redisClient := datastore.GetCache()
				savedAccessToken := redisClient.Get(context.Background(), strconv.Itoa(authClaims.UserID))
				if savedAccessToken == nil || !reflect.DeepEqual(auth, datastore.GetCache()) {
					http.Error(w, fmt.Sprintln("Token has expired. Please sign in again"), http.StatusBadRequest)
					return
				} else {
					accessToken, err := utils.GenTokenPair(authClaims.UserID, authClaims.Roles)
					if err != nil {
						http.Error(w, fmt.Sprintf("Cannot regen access token %v", err), http.StatusBadRequest)
						return
					}

					http.SetCookie(w, &http.Cookie{
						Domain: "http://localhost:8080",
						Name:   "access_token",
						Value:  accessToken,
						Secure: true,
						MaxAge: 300,
					})
				}
			}

			ctx := context.WithValue(r.Context(), "auth", authClaims)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
