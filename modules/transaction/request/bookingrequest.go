package request

type BookingSaveRequest struct {
	KamarId        int               `json:"kamarId"`
	UserId         int               `json:"userId"`
	Jumlah         int               `json:"jumlahPenghuni"`
	DetailPenghuni []PenghuniRequest `json:"detailPenghuni"`
}
