package request

import (
	"aplikasi-adakost-be/common"
	"errors"
)

type RegisterRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	FullName    string `json:"fullName"`
	NoHandphone string `json:"noHandphone"`
	Email       string `json:"email"`
	RoleId      int    `json:"roleid"`
}

func (l *RegisterRequest) ValidateLogin() (err error) {
	if common.IsEmptyField(l.Username) {
		return errors.New("username required")
	}

	if common.IsEmptyField(l.Password) {
		return errors.New("password required")
	}

	return
}
