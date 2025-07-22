package request

type UpdateKostRequest struct {
	NamaKost string `json:"namaKost"`
	Alamat   string `json:"alamat"`
	Pemilik  int    `json:"pemilik"`
	TypeKost string `json:"typeKost"`
}
