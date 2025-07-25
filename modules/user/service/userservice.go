package service

import (
	"aplikasi-adakost-be/middleware"
	"aplikasi-adakost-be/modules/user/model"
	"aplikasi-adakost-be/modules/user/repository"
	"aplikasi-adakost-be/modules/user/request"
	"aplikasi-adakost-be/modules/user/response"
	"errors"
	"fmt"
	"time"
)

const Bearer = "Bearer "

type UserService interface {
	SaveRegisterUser(req request.RegisterRequest) (response.SignUpResponse, error)
	Login(request request.Login) (model.UserLogin, error)
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
		CreatedBy: "SYSTEM",
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

func (s *userService) Login(request request.Login) (model.UserLogin, error) {

	if request.Username == "" {
		return model.UserLogin{}, errors.New("username tidak boleh kosong")
	}

	if request.Password == "" {
		return model.UserLogin{}, errors.New("password tidak boleh kosong")
	}

	userLogin := model.UserLogin{
		Username: request.Username,
		Password: request.Password,
	}

	userData, err := s.repo.Login(userLogin)
	if err != nil {
		return model.UserLogin{}, fmt.Errorf("login failed: %w", err)
	}

	// Pastikan ada roles dan ambil role pertama (atau sesuai kebutuhan)
	if len(userData.Roles) == 0 {
		return model.UserLogin{}, fmt.Errorf("user has no roles assigned")
	}

	roleName := userData.Roles[0].RoleName

	now := time.Now()
	expiry := now.Add(time.Hour * 1) // 1 jam dari sekarang

	token, err := middleware.GenerateJwtToken(userData.Id, userData.Username, roleName)
	if err != nil {
		return model.UserLogin{}, fmt.Errorf("token generation failed: %w", err)
	}

	userData.LoginAt = now
	userData.ExpiredAt = expiry
	userData.Token = Bearer + token

	return userData, nil
}
