package request

type PenghuniRequest struct {
	NamaPenghuni     string `json:"namaPenghuni"`
	NomorHp          int    `json:"nomorHp"`
	JenisKelamin     string `json:"jenisKelamin"`
	StatusPerkawinan string `json:"status"`
	NomorKtp         string `json:"nomorKtp"`
}
