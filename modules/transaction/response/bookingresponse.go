package response

type BookingResponse struct {
	Id             int                `json:"id"`
	NamaKost       string             `json:"namaKost"`
	TypeKost       string             `json:"typeKost"`
	JumlahPenghuni int                `json:"jumlahPenghuni"`
	StatusBooking  string             `json:"statusBooking"`
	DetailPenghuni []PenghuniResponse `json:"detailPenghuni"`
}
