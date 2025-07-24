package model

import (
	kamarmodel "aplikasi-adakost-be/modules/kamar/model"
	usermodel "aplikasi-adakost-be/modules/user/model"
	"time"
)

type Booking struct {
	Id             int              `json:"id"`
	User           usermodel.Users  `json:"users"`
	Kamar          kamarmodel.Kamar `json:"kamar"`
	TanggalMulai   time.Time        `json:"tanggalMulai"`
	TanggalAkhir   *time.Time       `json:"tanggalAkhir"`
	JumlahPenghuni int              `json:"jumlahPenghuni"`
	StatusBooking  string           `json:"statusBooking"`
	DetailPenghuni []BookingMember  `json:"detailPenghuni"`
	CreatedAt      time.Time        `json:"createdAt"`
	CreatedBy      string           `json:"createdBy"`
	ModifiedAt     time.Time        `json:"modifiedAt"`
	Modifiedby     string           `json:"modifiedBy"`
}
