package utils

import "github.com/minhtuhcmus/nbh-mono-be/constant"

func IsAdmin(roles []string) bool {
	for _, role := range roles {
		if role == constant.ADMIN {
			return true
		}
	}
	return false
}
