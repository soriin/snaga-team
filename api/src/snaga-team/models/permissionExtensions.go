package models

import (
	"snaga-team/config"
)

func (user *User) IsSystemAdmin() bool {
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
