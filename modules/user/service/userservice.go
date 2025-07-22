package service

import (
	"aplikasi-adakost-be/modules/user/model"
	"aplikasi-adakost-be/modules/user/repository"
	"aplikasi-adakost-be/modules/user/request"
	"aplikasi-adakost-be/modules/user/response"
	"errors"
	"fmt"
	"time"
)

type UserService interface {
	SaveRegisterUser(req request.RegisterRequest) (response.SignUpResponse, error)
}

type userService struct {
	repo repository.UsersRepository
}

func NewUserService(repo repository.UsersRepository) UserService {
	return &userService{repo: repo}
}

func (u *userService) SaveRegisterUser(req request.RegisterRequest) (response.SignUpResponse, error) {
	if req.Username == "" {
		return response.SignUpResponse{}, errors.New("username tidak boleh kosong")
	}

	if req.Password == "" {
		return response.SignUpResponse{}, errors.New("password tidak boleh kosong")
	}

	if len(req.Password) > 16 || len(req.Password) < 6 {
		return response.SignUpResponse{}, fmt.Errorf("panjang password tidak boleh %d", len(req.Password))
	}

	user := model.Users{
		Username:  req.Username,
		Password:  req.Password,
		RoleId:    req.RoleId,
		CreatedAt: time.Now(),
		CreatedBy: "Admin",
	}

	saveUser, err := u.repo.Register(user)
	if err != nil {
		return response.SignUpResponse{}, err
	}

	resp := response.SignUpResponse{
		Username: saveUser.Username,
		Password: saveUser.Password,
	}

	return resp, nil
}
