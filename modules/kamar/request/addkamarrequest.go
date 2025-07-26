package request

type KamarRequest struct {
	NamaKamar  string `json:"namaKamar"`
	HargaKamar int    `json:"hargaKamar"`
	KostId     int    `json:"kostId"`
}
