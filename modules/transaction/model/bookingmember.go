package model

import "time"

type BookingMember struct {
	Id            int       `json:"id"`
	Booking       Booking   `json:"booking"`
	NamaPenghuni  string    `json:"namaPenghuni"`
	NomorHp       int       `json:"nomorHp"`
	JenisKelamin  string    `json:"jenisKelamin"`
	MaritalStatus string    `json:"statusKtp"`
	NomorKtp      string    `json:"nomorKtp"`
	CreatedAt     time.Time `json:"createdAt"`
	CreatedBy     string    `json:"createdBy"`
	ModifiedAt    time.Time `json:"modifiedAt"`
	Modifiedby    string    `json:"modifiedBy"`
}
