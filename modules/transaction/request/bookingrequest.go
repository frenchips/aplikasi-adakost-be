package request

type BookingSaveRequest struct {
	KamarId        int               `json:"kamarId"`
	Jumlah         int               `json:"jumlahPenghuni"`
	DetailPenghuni []PenghuniRequest `json:"detailPenghuni"`
}
