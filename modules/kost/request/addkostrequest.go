package request

type AddKostRequest struct {
	NamaKost string `json:"namaKost"`
	Alamat   string `json:"alamat"`
	TypeKost string `json:"typeKost"`
}
