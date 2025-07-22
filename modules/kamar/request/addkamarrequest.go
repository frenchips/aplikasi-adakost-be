package request

type KamarRequest struct {
	NomorKamar  string `json:"nomorKamar"`
	HargaKamar  int    `json:"hargaKamar"`
	StatusKamar string `json:"statusKamar"`
	KostId      int    `json:"kostId"`
}
