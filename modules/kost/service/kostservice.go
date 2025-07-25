package service

import (
	kostmodel "aplikasi-adakost-be/modules/kost/model"
	"aplikasi-adakost-be/modules/kost/repository"
	"aplikasi-adakost-be/modules/kost/request"
	"aplikasi-adakost-be/modules/kost/response"
	usermodel "aplikasi-adakost-be/modules/user/model"

	"errors"
	"time"
)

type KostService interface {
	InsertKost(req request.AddKostRequest) (response.KostResponse, error)
	UpdateKost(req request.UpdateKostRequest, id int) (response.KostResponse, error)
	GetAllKost(kost response.ViewKostResponse) (result []response.ViewKostResponse, err error)
	DeleteKost(id int) error
	GetKostKamar() (result []response.KamarKostReponse, err error)
}

type kostService struct {
	repo repository.KostRepository
}

func NewKostService(repo repository.KostRepository) KostService {
	return &kostService{repo: repo}
}

func (u *kostService) InsertKost(req request.AddKostRequest) (response.KostResponse, error) {
	if req.NamaKost == "" {
		return response.KostResponse{}, errors.New("username tidak boleh kosong")
	}

	if req.Alamat == "" {
		return response.KostResponse{}, errors.New("password tidak boleh kosong")
	}

	// if len(req.Password) > 16 || len(req.Password) < 6 {
	// 	return response.SignUpResponse{}, fmt.Errorf("panjang password tidak boleh %d", len(req.Password))
	// }

	kost := kostmodel.Kost{
		NamaKost: req.NamaKost,
		Alamat:   req.Alamat,
		Users: usermodel.Users{
			Id: req.Pemilik,
		},
		TypeKost:  req.TypeKost,
		CreatedAt: time.Now(),
		CreatedBy: "Admin",
	}

	saveUser, err := u.repo.InserKost(kost)
	if err != nil {
		return response.KostResponse{}, err
	}

	resp := response.KostResponse{
		NamaKost: saveUser.NamaKost,
		Alamat:   saveUser.Alamat,
		Pemilik:  saveUser.Users.Username,
		TypeKost: saveUser.TypeKost,
	}

	return resp, nil
}

func (u *kostService) UpdateKost(req request.UpdateKostRequest, id int) (response.KostResponse, error) {
	if req.NamaKost == "" {
		return response.KostResponse{}, errors.New("username tidak boleh kosong")
	}

	if req.Alamat == "" {
		return response.KostResponse{}, errors.New("password tidak boleh kosong")
	}

	// if len(req.Password) > 16 || len(req.Password) < 6 {
	// 	return response.SignUpResponse{}, fmt.Errorf("panjang password tidak boleh %d", len(req.Password))
	// }

	kost := kostmodel.Kost{
		NamaKost: req.NamaKost,
		Alamat:   req.Alamat,
		Users: usermodel.Users{
			Id: req.Pemilik,
		},
		TypeKost:  req.TypeKost,
		CreatedAt: time.Now(),
		CreatedBy: "Admin",
	}

	saveUser, err := u.repo.UpdateKost(kost, id)
	if err != nil {
		return response.KostResponse{}, err
	}

	resp := response.KostResponse{
		NamaKost: saveUser.NamaKost,
		Alamat:   saveUser.Alamat,
		Pemilik:  saveUser.Users.Username,
		TypeKost: saveUser.TypeKost,
	}

	return resp, nil
}

func (c *kostService) GetAllKost(kost response.ViewKostResponse) (result []response.ViewKostResponse, err error) {
	return c.repo.GetAllKost(kost)
}

func (c *kostService) DeleteKost(id int) error {
	return c.repo.DeleteKost(id)
}

func (c *kostService) GetKostKamar() (result []response.KamarKostReponse, err error) {
	return c.repo.GetKostKamar()
}
