package response

// import kostresponse "aplikasi-adakost-be/modules/kost/response"

type KamarResponse struct {
	NomorKamar  string `json:"nomorKamar"`
	HargaKamar  int    `json:"hargaKamar"`
	StatusKamar string `json:"statusKamar"`
	// Kost        kostresponse.KostResponse `json:"kost"`
}
