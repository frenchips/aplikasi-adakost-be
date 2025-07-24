package response

type BookingSaveResponse struct {
	NamaKost       string             `json:"namaKost"`
	Jumlah         int                `json:"jumlahPenghuni"`
	StatusBooking  string             `json:"statusBooking"`
	DetailPenghuni []PenghuniResponse `json:"detailPenghuni"`
}
