package service

import (
	kamarmodel "aplikasi-adakost-be/modules/kamar/model"
	kamarrepository "aplikasi-adakost-be/modules/kamar/repository"
	bookingmodel "aplikasi-adakost-be/modules/transaction/model"
	bookingrepository "aplikasi-adakost-be/modules/transaction/repository"
	bookingrequest "aplikasi-adakost-be/modules/transaction/request"
	"aplikasi-adakost-be/modules/transaction/response"
	bookingresponse "aplikasi-adakost-be/modules/transaction/response"
	usermodel "aplikasi-adakost-be/modules/user/model"
	"fmt"
	"time"
)

const (
	StatusConfirm = "Confirmed"
	StatusTerisi  = "Terisi"
	StatusCancel  = "Canceled"
)

type TransactionService interface {
	SaveOrderBooking(request bookingrequest.BookingSaveRequest, id int) (bookingresponse.BookingSaveResponse, error)
	CancelOrderBooking(id int) error
	GetDetailBooking(id int) ([]response.BookingResponse, error)
	GetDetailUserBooking(id int) ([]response.BookingResponse, error)
}

type transactionService struct {
	repo      bookingrepository.TransactionRepository
	kamarRepo kamarrepository.KamarRepository
}

func NewTransactionService(repo bookingrepository.TransactionRepository, kamarRepo kamarrepository.KamarRepository) TransactionService {
	return &transactionService{repo: repo, kamarRepo: kamarRepo}
}

func (t *transactionService) SaveOrderBooking(request bookingrequest.BookingSaveRequest, id int) (bookingresponse.BookingSaveResponse, error) {

	if request.Jumlah > 2 {
		return bookingresponse.BookingSaveResponse{}, fmt.Errorf("maksimal penghuni 2 tidak boleh lebih %d", request.Jumlah)
	}

	booking := bookingmodel.Booking{
		User: usermodel.Users{
			Id: request.UserId,
		},
		Kamar: kamarmodel.Kamar{
			Id: request.KamarId,
		},
		TanggalMulai:   time.Now(),
		TanggalAkhir:   nil,
		JumlahPenghuni: request.Jumlah,
		StatusBooking:  StatusConfirm,
		CreatedAt:      time.Now(),
		CreatedBy:      "Admin",
		DetailPenghuni: func() []bookingmodel.BookingMember {
			var detailPenghuni []bookingmodel.BookingMember

			for _, penghuniReq := range request.DetailPenghuni {
				detailPenghuni = append(detailPenghuni, bookingmodel.BookingMember{
					NamaPenghuni:  penghuniReq.NamaPenghuni,
					JenisKelamin:  penghuniReq.JenisKelamin,
					NomorHp:       penghuniReq.NomorHp,
					MaritalStatus: penghuniReq.StatusPerkawinan,
					NomorKtp:      penghuniReq.NomorKtp,
					CreatedAt:     time.Now(),
					CreatedBy:     "Admin",
				})
			}
			return detailPenghuni
		}(),
	}

	saveOrder, err := t.repo.SaveBooking(booking, id)
	if err != nil {
		return bookingresponse.BookingSaveResponse{}, err
	}

	_, err = t.kamarRepo.UpdateKostKamar(kamarmodel.Kamar{
		Id:          request.KamarId,
		StatusKamar: StatusTerisi,
	}, request.KamarId)

	if err != nil {
		return bookingresponse.BookingSaveResponse{}, fmt.Errorf("gagal update status kamar: %v", err)
	}

	resp := bookingresponse.BookingSaveResponse{
		NamaKost:      saveOrder.Kamar.Kost.NamaKost,
		Jumlah:        saveOrder.JumlahPenghuni,
		StatusBooking: saveOrder.StatusBooking,
		DetailPenghuni: func() []bookingresponse.PenghuniResponse {
			var penghuniResponses []bookingresponse.PenghuniResponse
			for _, penghuni := range saveOrder.DetailPenghuni {
				penghuniResponses = append(penghuniResponses, bookingresponse.PenghuniResponse{
					NamaPenghuni:     penghuni.NamaPenghuni,
					NomorHp:          penghuni.NomorHp,
					JenisKelamin:     penghuni.JenisKelamin,
					StatusPerkawinan: penghuni.MaritalStatus,
					NomorKtp:         penghuni.NomorKtp,
				})
			}
			return penghuniResponses
		}(),
	}

	return resp, nil
}

func (t *transactionService) CancelOrderBooking(id int) error {

	existing, err := t.repo.FindBookingById(id)
	if err != nil {
		return fmt.Errorf("booking tidak ditemukan")
	}

	if existing.StatusBooking == "Cancelled" {
		return fmt.Errorf("booking sudah dibatalkan sebelumnya")
	}

	err = t.repo.UpdateBookingStatus(id, StatusCancel, "Admin") // atau ambil dari session user
	if err != nil {
		return fmt.Errorf("gagal membatalkan booking: %v", err)
	}

	err = t.kamarRepo.UpdateKamarStatus(existing.Kamar.Id, "Belum terisi")
	if err != nil {
		return fmt.Errorf("gagal update status kamar: %v", err)
	}

	return nil
}

func (t *transactionService) GetDetailBooking(id int) ([]response.BookingResponse, error) {
	result, err := t.repo.GetDetailBooking(id)
	if err != nil {
		return nil, fmt.Errorf("gagal menampilkan detail booking: %v", err)
	}

	return result, nil
}

func (t *transactionService) GetDetailUserBooking(id int) ([]response.BookingResponse, error) {
	result, err := t.repo.GetDetailUsersBooking(id)
	if err != nil {
		return nil, fmt.Errorf("gagal menampilkan detail booking: %v", err)
	}

	return result, nil
}
