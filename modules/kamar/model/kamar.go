package model

import (
	"aplikasi-adakost-be/modules/kost/model"
	"time"
)

type Kamar struct {
	Id          int        `json:"id"`
	Kost        model.Kost `json:"kost"`
	NamaKamar   string     `json:"namaKamar"`
	HargaKamar  int        `json:"hargaKamar"`
	StatusKamar string     `json:"statusKamar"`
	CreatedAt   time.Time  `json:"createdAt"`
	CreatedBy   string     `json:"createdBy"`
	ModifiedAt  time.Time  `json:"modifiedAt"`
	ModifiedBy  string     `json:"modifiedBy"`
}
