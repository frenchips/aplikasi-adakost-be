package response

import "aplikasi-adakost-be/modules/kamar/response"

type KamarKostReponse struct {
	Id          int                         `json:"id"`
	NamaKost    string                      `json:"namaKost"`
	Alamat      string                      `json:"alamat"`
	TypeKost    string                      `json:"typeKost"`
	SisaKamar   int                         `json:"sisaKamar"`
	DetailKamar []response.GetKamarResponse `json:"detailKamar"`
}
