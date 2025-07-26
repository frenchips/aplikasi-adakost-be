package repository

import (
	kamarmodel "aplikasi-adakost-be/modules/kamar/model"
	"aplikasi-adakost-be/modules/kamar/response"
	"database/sql"
	"time"
)

type KamarRepository interface {
	InsertKamar(kamar kamarmodel.Kamar) (kamarmodel.Kamar, error)
	UpdateKamar(kamar kamarmodel.Kamar, id int) (kamarmodel.Kamar, error)
	DeleteKamar(id int) error
	GetAllKost(kamar response.GetKamarResponse, userId int) (result []response.GetKamarResponse, err error)
	UpdateKostKamar(kamar kamarmodel.Kamar, id int) (kamarmodel.Kamar, error)
	UpdateKamarStatus(kamarId int, status string) error
}

type kamarRepository struct {
	db *sql.DB
}

func NewKamarRepository(db *sql.DB) KamarRepository {
	return &kamarRepository{db: db}
}

func (k *kamarRepository) InsertKamar(kamar kamarmodel.Kamar) (kamarmodel.Kamar, error) {
	now := time.Now()
	createBy := "Admin"

	sql := "INSERT INTO adk_kamar(kost_id, nama_kamar, harga_per_bulan, status_kamar, created_at, created_by) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	errs := k.db.QueryRow(sql, &kamar.Kost.Id, &kamar.NamaKamar, &kamar.HargaKamar, &kamar.StatusKamar, &now, &createBy).Scan(&kamar.Id)
	if errs != nil {
		panic(errs)
	}

	var nameKost, alamat, namaPemilik, typeKost string
	errs = k.db.QueryRow("SELECT ak.nama_kost, ak.alamat, au.username, ak.type_kost FROM adk_kost ak join adk_users au on ak.pemilik_id = au.id WHERE ak.id = $1", kamar.Kost.Id).Scan(&nameKost, &alamat, &namaPemilik, &typeKost)
	if errs != nil {
		return kamar, errs
	}

	kamar.Kost.NamaKost = nameKost
	kamar.Kost.Users.Username = namaPemilik
	kamar.Kost.TypeKost = typeKost
	return kamar, nil
}

func (k *kamarRepository) UpdateKamar(kamar kamarmodel.Kamar, id int) (kamarmodel.Kamar, error) {
	now := time.Now()
	modifiedBy := "Admin"

	sql := "UPDATE adk_kamar SET kost_id = $1,  harga_per_bulan = $2, status_kamar = $3, modified_at = $4,  modified_by = $5 WHERE id = $6 RETURNING id"
	errs := k.db.QueryRow(sql, &kamar.Kost.Id, &kamar.HargaKamar, &kamar.StatusKamar, &now, &modifiedBy, id).Scan(&kamar.Id)
	if errs != nil {
		panic(errs)
	}

	kamar.ModifiedBy = modifiedBy
	kamar.ModifiedAt = now

	var nameKost, alamat, namaPemilik, typeKost string
	errs = k.db.QueryRow("SELECT ak.nama_kost, ak.alamat, au.username, ak.type_kost FROM adk_kost ak join adk_users au on ak.pemilik_id = au.id WHERE ak.id = $1", kamar.Kost.Id).Scan(&nameKost, &alamat, &namaPemilik, &typeKost)
	if errs != nil {
		return kamar, errs
	}

	kamar.Kost.NamaKost = nameKost
	kamar.Kost.Users.Username = namaPemilik
	kamar.Kost.TypeKost = typeKost
	return kamar, nil
}

func (k *kamarRepository) UpdateKostKamar(kamar kamarmodel.Kamar, id int) (kamarmodel.Kamar, error) {
	now := time.Now()
	modifiedBy := "Admin"

	sql := "UPDATE adk_kamar SET  status_kamar = $1, modified_at = $2,  modified_by = $3 WHERE id = $4 RETURNING id"
	errs := k.db.QueryRow(sql, &kamar.StatusKamar, &now, &modifiedBy, id).Scan(&kamar.Id)
	if errs != nil {
		panic(errs)
	}

	kamar.ModifiedBy = modifiedBy
	kamar.ModifiedAt = now

	return kamar, nil
}

func (k *kamarRepository) UpdateKamarStatus(kamarId int, status string) error {
	now := time.Now()
	modifiedBy := "Admin"
	_, err := k.db.Exec(
		"UPDATE adk_kamar SET status_kamar = $1, modified_at = $2, modified_by = $3 WHERE id = $4",
		status, now, modifiedBy, kamarId,
	)
	return err
}

func (k *kamarRepository) DeleteKamar(id int) error {

	sql := "DELETE FROM adk_kamar WHERE id = $1"
	_, errs := k.db.Exec(sql, id)
	if errs != nil {
		panic(errs)
	}

	return nil
}

func (k *kamarRepository) GetAllKost(kamar response.GetKamarResponse, userId int) (result []response.GetKamarResponse, err error) {
	sql := `select 
				ak.nama_kamar, 
				ak.harga_per_bulan, 
				ak.status_kamar  
			from adk_kamar ak  
			join adk_kost aks on ak.kost_id = aks.id
			join adk_users au on aks.pemilik_id = au.id
			where aks.pemilik_id = $1
			ORDER BY ak.id ASC`
	rows, err := k.db.Query(sql, userId)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var kamar response.GetKamarResponse

		err = rows.Scan(&kamar.NomorKamar, &kamar.HargaKamar, &kamar.StatusKamar)
		if err != nil {
			return
		}

		result = append(result, kamar)
	}

	return
}
