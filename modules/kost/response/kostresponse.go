package response

type ViewKostResponse struct {
	NamaKost string `json:"namaKost"`
	Alamat   string `json:"alamat"`
	Pemilik  string `json:"pemilik"`
	TypeKost string `json:"typeKost"`
}
