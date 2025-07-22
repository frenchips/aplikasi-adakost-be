package request

type UpdateKamarRequest struct {
	HargaKamar  int    `json:"hargaKamar"`
	StatusKamar string `json:"statusKamar"`
	KostId      int    `json:"kostId"`
}
