package repository

import (
	"aplikasi-adakost-be/modules/kost/model"
	"aplikasi-adakost-be/modules/kost/response"
	"database/sql"
	"time"
)

type KostRepository interface {
	InserKost(kost model.Kost) (model.Kost, error)
	UpdateKost(kost model.Kost, id int) (model.Kost, error)
	GetAllKost(kost response.ViewKostResponse) (result []response.ViewKostResponse, err error)
	DeleteKost(id int) error
}

type kostRepository struct {
	db *sql.DB
}

func NewKostService(db *sql.DB) KostRepository {
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

func (c *kostRepository) GetAllKost(kost response.ViewKostResponse) (result []response.ViewKostResponse, err error) {
	sql := "select ak.nama_kost, au.username, ak.alamat, ak.type_kost  from adk_kost ak join adk_users au on ak.pemilik_id = au.id  ORDER BY ak.id ASC"
	rows, err := c.db.Query(sql)
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
