package model

import (
	"aplikasi-adakost-be/modules/user/model"
	"time"
)

type Kost struct {
	Id         int         `json:"id"`
	NamaKost   string      `json:"namaKost"`
	Alamat     string      `json:"alamat"`
	Users      model.Users `json:"users"`
	TypeKost   string      `json:"typeKost"`
	CreatedAt  time.Time   `json:"createdAt"`
	CreatedBy  string      `json:"createdBy"`
	ModifiedAt time.Time   `json:"modifiedAt"`
	ModifiedBy string      `json:"modifiedBy"`
}
