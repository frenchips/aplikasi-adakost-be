package repository

import (
	kamarrepo "aplikasi-adakost-be/modules/kamar/repository"
	kamarresponse "aplikasi-adakost-be/modules/kamar/response"
	"aplikasi-adakost-be/modules/kost/model"
	"aplikasi-adakost-be/modules/kost/response"
	"database/sql"
	"time"
)

type KostRepository interface {
	InserKost(kost model.Kost) (model.Kost, error)
	UpdateKost(kost model.Kost, id int) (model.Kost, error)
	GetAllKost(kost response.ViewKostResponse, userId int) (result []response.ViewKostResponse, err error)
	DeleteKost(id int) error
	GetKostKamar() (result []response.KamarKostReponse, err error)
	GetKamarByKost(id int) (result []kamarresponse.GetKamarResponse, err error)
}

type kostRepository struct {
	db *sql.DB
}

func NewKostRepository(db *sql.DB, kamarrepo kamarrepo.KamarRepository) KostRepository {
	return &kostRepository{db: db}
}

func (k *kostRepository) InserKost(kost model.Kost) (model.Kost, error) {
	sql := "INSERT INTO adk_kost(nama_kost, alamat, pemilik_id, type_kost, created_at, created_by) VALUES($1, $2, $3, $4, $5, $6) RETURNING id"
	errs := k.db.QueryRow(sql, &kost.NamaKost, &kost.Alamat, &kost.Users.Id, &kost.TypeKost, &kost.CreatedAt, &kost.CreatedBy).Scan(&kost.Id)
	if errs != nil {
		panic(errs)
	}

	var name string
	errs = k.db.QueryRow("SELECT username FROM adk_users WHERE id = $1", kost.Users.Id).Scan(&name)
	if errs != nil {
		return kost, errs
	}

	kost.Users.Username = name

	return kost, nil
}

func (k *kostRepository) UpdateKost(kost model.Kost, id int) (model.Kost, error) {

	now := time.Now()
	updateBy := "admin"

	sql := "UPDATE adk_kost SET nama_kost = $1, alamat = $2, pemilik_id = $3, type_kost = $4, modified_at = $5, modified_by = $6 WHERE id = $7 RETURNING id"
	errs := k.db.QueryRow(sql, kost.NamaKost, kost.Alamat, kost.Users.Id, kost.TypeKost, kost.ModifiedAt, kost.ModifiedBy, id).Scan(&kost.Id)
	if errs != nil {
		panic(errs)
	}

	var name string
	errs = k.db.QueryRow("SELECT username FROM adk_users WHERE id = $1", kost.Users.Id).Scan(&name)
	if errs != nil {
		return kost, errs
	}

	kost.Users.Username = name

	kost.ModifiedAt = now
	kost.ModifiedBy = updateBy

	return kost, nil
}

func (c *kostRepository) GetAllKost(kost response.ViewKostResponse, userId int) (result []response.ViewKostResponse, err error) {
	sql := "select ak.nama_kost, au.username, ak.alamat, ak.type_kost  from adk_kost ak join adk_users au on ak.pemilik_id = au.id WHERE ak.pemilik_id = $1  ORDER BY ak.id ASC"
	rows, err := c.db.Query(sql, userId)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var kost response.ViewKostResponse

		err = rows.Scan(&kost.NamaKost, &kost.Alamat, &kost.Pemilik, &kost.TypeKost)
		if err != nil {
			return
		}

		result = append(result, kost)
	}

	return
}

func (c *kostRepository) DeleteKost(id int) error {

	sql := "DELETE FROM adk_kost WHERE id = $1"
	_, errs := c.db.Exec(sql, id)
	if errs != nil {
		return errs
	}
	return nil
}

func (c *kostRepository) GetKostKamar() (result []response.KamarKostReponse, err error) {
	sql := `
		SELECT
			aks.id,
			aks.nama_kost,
			aks.alamat,
			aks.type_kost,
			count(ak.status_kamar) as sisa_kamar
		FROM adk_kost aks 
		join adk_kamar ak on aks.id = ak.kost_id
		join adk_users au on aks.pemilik_id = au.id 
		where ak.status_kamar = 'Belum terisi' 
		GROUP BY aks.id, aks.nama_kost, aks.type_kost
	`
	rows, err := c.db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var kost response.KamarKostReponse

		err = rows.Scan(&kost.Id, &kost.NamaKost, &kost.Alamat, &kost.TypeKost, &kost.SisaKamar)
		if err != nil {
			return
		}

		kost.DetailKamar, err = c.GetKamarByKost(kost.Id)
		if err != nil {
			return
		}

		result = append(result, kost)
	}

	return
}

func (c *kostRepository) GetKamarByKost(id int) (result []kamarresponse.GetKamarResponse, err error) {
	sql := `
		SELECT
			ak.nama_kamar,
			ak.harga_per_bulan,
			ak.status_kamar
		FROM adk_kamar ak
		JOIN adk_kost aks ON ak.kost_id = aks.id
		WHERE aks.id = $1
	`
	rows, err := c.db.Query(sql, id)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var kamar kamarresponse.GetKamarResponse

		err = rows.Scan(&kamar.NomorKamar, &kamar.HargaKamar, &kamar.StatusKamar)
		if err != nil {
			return
		}

		result = append(result, kamar)
	}

	return
}
