package service

import (
	kamarmodel "aplikasi-adakost-be/modules/kamar/model"
	"aplikasi-adakost-be/modules/kamar/repository"
	kamarrequest "aplikasi-adakost-be/modules/kamar/request"
	kamarresponse "aplikasi-adakost-be/modules/kamar/response"
	kostmodel "aplikasi-adakost-be/modules/kost/model"
	"time"
)

type KamarService interface {
	InsertKamar(req kamarrequest.KamarRequest, username string) (kamarresponse.KamarResponse, error)
	UpdateKamar(req kamarrequest.UpdateKamarRequest, id int, username string) (kamarresponse.KamarResponse, error)
	DeleteKamar(id int) error
	GetAllKamar(kamar kamarresponse.GetKamarResponse) (result []kamarresponse.GetKamarResponse, err error)
}

type kamarService struct {
	repo repository.KamarRepository
}

func NewKamarService(repo repository.KamarRepository) KamarService {
	return &kamarService{repo: repo}
}

func (k *kamarService) InsertKamar(req kamarrequest.KamarRequest, username string) (kamarresponse.KamarResponse, error) {

	kamar := kamarmodel.Kamar{
		NamaKamar:   req.NamaKamar,
		HargaKamar:  req.HargaKamar,
		StatusKamar: req.StatusKamar,
		Kost: kostmodel.Kost{
			Id: req.KostId,
		},
		CreatedAt: time.Now(),
		CreatedBy: username,
	}

	saveKamar, err := k.repo.InsertKamar(kamar)
	if err != nil {
		return kamarresponse.KamarResponse{}, err
	}

	resp := kamarresponse.KamarResponse{
		NomorKamar:  saveKamar.NamaKamar,
		HargaKamar:  saveKamar.HargaKamar,
		StatusKamar: saveKamar.StatusKamar,
		// Kost: kostresponse.KostResponse{
		// 	NamaKost: saveKamar.Kost.NamaKost,
		// 	Pemilik:  saveKamar.Kost.Users.Username,
		// },
	}

	return resp, nil
}

func (u *kamarService) UpdateKamar(req kamarrequest.UpdateKamarRequest, id int, username string) (kamarresponse.KamarResponse, error) {
	// if req.NamaKost == "" {
	// 	return response.KostResponse{}, errors.New("username tidak boleh kosong")
	// }

	// if req.Alamat == "" {
	// 	return response.KostResponse{}, errors.New("password tidak boleh kosong")
	// }

	// if len(req.Password) > 16 || len(req.Password) < 6 {
	// 	return response.SignUpResponse{}, fmt.Errorf("panjang password tidak boleh %d", len(req.Password))
	// }

	kamar := kamarmodel.Kamar{
		HargaKamar:  req.HargaKamar,
		StatusKamar: req.StatusKamar,
		Kost: kostmodel.Kost{
			Id: req.KostId,
		},
		ModifiedAt: time.Now(),
		ModifiedBy: username,
	}

	saveUser, err := u.repo.UpdateKamar(kamar, id)
	if err != nil {
		return kamarresponse.KamarResponse{}, err
	}

	resp := kamarresponse.KamarResponse{
		NomorKamar:  saveUser.NamaKamar,
		HargaKamar:  saveUser.HargaKamar,
		StatusKamar: saveUser.StatusKamar,
		// Kost: kostresponse.KostResponse{
		// 	NamaKost: saveUser.Kost.NamaKost,
		// },
	}

	return resp, nil
}

func (k *kamarService) DeleteKamar(id int) error {
	return k.repo.DeleteKamar(id)
}

func (k *kamarService) GetAllKamar(kamar kamarresponse.GetKamarResponse) (result []kamarresponse.GetKamarResponse, err error) {
	return k.repo.GetAllKost(kamar)
}
