package models

import (
	"snaga-team/config"
)

func (user *User) IsAdmin() bool {
	if user.Email == config.MASTER_EMAIL {
		return true
	}

	for _, a := range user.Groups {
		if a == "admin" {
			return true
		}
	}
	return false
}
