package model

import "time"

type Users struct {
	Id         int       `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	FullName   string    `json:"fullName"`
	NoHp       string    `json:"noHp"`
	Email      string    `json:"email"`
	RoleId     int       `json:"roleId"`
	CreatedAt  time.Time `json:"createdAt"`
	CreatedBy  string    `json:"createdBy"`
	ModifiedAt time.Time `json:"modifiedAt"`
	Modifiedby string    `json:"modifiedBy"`
}
