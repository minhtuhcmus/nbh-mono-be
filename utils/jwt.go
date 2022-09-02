package utils

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/minhtuhcmus/nbh-mono-be/database/datastore"
	"strconv"
	"time"
)

type claims struct {
	UserID      int      `json:"userID"`
	Permissions []string `json:"permissions"`
	jwt.StandardClaims
}

var salt = []byte("n4aban40a_very_secret_key")

func JwtGenerate(userID int, permissions []string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims{
		UserID:      userID,
		Permissions: permissions,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	token, err := t.SignedString(salt)
	if err != nil {
		return "", fmt.Errorf("error when sign token %v", err)
	}

	return token, nil
}

func JwtValidate(token string) (*claims, error) {
	var auth claims
	parsedToken, err := jwt.ParseWithClaims(token, &auth, func(token *jwt.Token) (interface{}, error) {
		return salt, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, fmt.Errorf("error when parse token %v", err)
		}
		return nil, fmt.Errorf("unexpected error %v", err)
	}

	if !parsedToken.Valid {
		return nil, fmt.Errorf("unauthorized %v", err)
	}

	return &auth, nil
}

func GenTokenPair(userID int, permissions []string) (string, error) {
	accessToken, err := JwtGenerate(userID, permissions)
	if err != nil {
		return "", err
	}

	datastore.GetCache().Set(context.Background(), strconv.Itoa(userID), accessToken, time.Hour*24*3)

	return accessToken, nil
}
