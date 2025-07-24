package request

type KamarRequest struct {
	NamaKamar   string `json:"namaKamar"`
	HargaKamar  int    `json:"hargaKamar"`
	StatusKamar string `json:"statusKamar"`
	KostId      int    `json:"kostId"`
}
