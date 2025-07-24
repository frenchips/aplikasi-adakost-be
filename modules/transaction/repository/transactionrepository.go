package repository

import (
	bookingmodel "aplikasi-adakost-be/modules/transaction/model"
	"database/sql"
	"fmt"
	"time"
)

type TransactionRepository interface {
	SaveBooking(ordermodel bookingmodel.Booking, id int) (bookingmodel.Booking, error)
	FindBookingById(id int) (bookingmodel.Booking, error)
	UpdateBookingStatus(id int, status string, modifiedBy string) error
}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (t *transactionRepository) SaveBooking(ordermodel bookingmodel.Booking, id int) (bookingmodel.Booking, error) {

	now := time.Now()
	by := "Admin"

	var kostId int
	sqlKost := `SELECT id FROM adk_kost WHERE id = $1`
	errs := t.db.QueryRow(sqlKost, id).Scan(&kostId)

	if errs != nil {
		if errs == sql.ErrNoRows {
			return ordermodel, fmt.Errorf("kost dengan ID %d tidak ditemukan", kostId)
		}
		return ordermodel, fmt.Errorf("gagal validasi kost: %v", errs)
	}

	sql := `INSERT INTO adk_booking(user_id, kamar_id, tanggal_mulai, tanggal_akhir, jumlah_penghuni, status_booking, created_at, created_by)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	errs = t.db.QueryRow(sql, &ordermodel.User.Id, &ordermodel.Kamar.Id, &now, &ordermodel.TanggalAkhir, &ordermodel.JumlahPenghuni, &ordermodel.StatusBooking, &now, by).Scan(&ordermodel.Id)

	if errs != nil {
		panic(errs)
	}

	for _, penghuni := range ordermodel.DetailPenghuni {
		penghuni.Booking.Id = ordermodel.Id
		_, err := t.saveBookingModel(penghuni)
		if err != nil {
			return ordermodel, err
		}
	}

	var namaKost string
	sqlGetKostName := `
	SELECT k.nama_kost 
	FROM adk_kost k
	JOIN adk_kamar km ON km.kost_id = k.id
	WHERE km.id = $1
`
	err := t.db.QueryRow(sqlGetKostName, ordermodel.Kamar.Id).Scan(&namaKost)
	if err != nil {
		return ordermodel, fmt.Errorf("gagal ambil nama kost: %v", err)
	}
	ordermodel.Kamar.Kost.NamaKost = namaKost

	return ordermodel, nil
}

func (t *transactionRepository) saveBookingModel(ordermodel bookingmodel.BookingMember) (bookingmodel.BookingMember, error) {

	now := time.Now()
	by := "Admin"

	sql := `INSERT INTO adk_booking_member(booking_id, nama_penghuni, jenis_kelamin, nomor_handphone, status_perkawinan, ktp_penghuni, created_at, created_by)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	errs := t.db.QueryRow(sql, &ordermodel.Booking.Id, &ordermodel.NamaPenghuni, &now, &ordermodel.NomorHp, &ordermodel.MaritalStatus, &ordermodel.NomorKtp, &now, by).Scan(&ordermodel.Id)

	if errs != nil {
		panic(errs)
	}

	return ordermodel, nil
}

func (r *transactionRepository) FindBookingById(id int) (bookingmodel.Booking, error) {
	var booking bookingmodel.Booking
	query := "SELECT id, status_booking, kamar_id FROM adk_booking WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&booking.Id, &booking.StatusBooking, &booking.Kamar.Id)
	return booking, err
}

func (r *transactionRepository) UpdateBookingStatus(id int, status string, modifiedBy string) error {
	now := time.Now()
	query := "UPDATE adk_booking SET status_booking = $1, modified_at = $2, modified_by = $3 WHERE id = $4"
	_, err := r.db.Exec(query, status, now, modifiedBy, id)
	return err
}
