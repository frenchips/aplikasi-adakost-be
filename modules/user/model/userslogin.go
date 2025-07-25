package model

import "time"

type UserLogin struct {
	Id        int       `json:"userId"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	LoginAt   time.Time `json:"loginAt"`
	ExpiredAt time.Time `json:"expiredAt"`
	Roles     []Roles   `json:"roles"`
	Token     string    `json:"token"`
}
