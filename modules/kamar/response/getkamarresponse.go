package response

type GetKamarResponse struct {
	NomorKamar  string `json:"nomorKamar"`
	HargaKamar  int    `json:"hargaKamar"`
	StatusKamar string `json:"statusKamar"`
}
