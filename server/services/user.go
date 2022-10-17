package services

import (
	"final-project/server/controllers/view"
	"final-project/server/repositories"
	"final-project/server/repositories/models"
	"final-project/server/request"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo repositories.UserRepo
}

func NewUserService(userRepo repositories.UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Register(req *request.CreateUserRequest) (view.ResponseRegisterUser, error) {
	var user models.User
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return view.ResponseRegisterUser{}, err
	}

	user.Username = req.Username
	user.Email = req.Email
	user.Password = string(hash)
	user.Age = req.Age
	user.CreatedAt = time.Now()

	userId, err := s.userRepo.Create(&user)

	if err != nil {
		return view.ResponseRegisterUser{}, err
	}

	return view.ResponseRegisterUser{
		Id:       userId,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}, nil
}

func (s *UserService) Login(req *request.UserLoginRequest) (string, error) {
	data, err := s.userRepo.FindByEmail(req.Email)

	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(req.Password))

	if err != nil {
		return "", err
	}

	return req.Email, nil
}

func (s *UserService) Update(id int, req *request.UpdateUserRequest) (view.ResponseUpdateUser, error) {
	var user models.User
	user.Username = req.Username
	user.Email = req.Email

	data, err := s.userRepo.UpdateById(id, &user)

	if err != nil {
		return view.ResponseUpdateUser{}, err
	}

	return view.ResponseUpdateUser{
		Id:        data.Id,
		Username:  data.Username,
		Email:     data.Email,
		Age:       data.Age,
		UpdatedAt: data.UpdatedAt,
	}, nil
}

func (s *UserService) Delete(email string) (view.ResponseDeleteUser, error) {
	err := s.userRepo.DeleteByEmail(email)

	if err != nil {
		return view.ResponseDeleteUser{}, err
	}

	return view.ResponseDeleteUser{
		Message: "Your account has been successfully deleted",
	}, nil
}

func (s *UserService) GetUserIdByEmail(email string) (int, error) {
	data, err := s.userRepo.FindByEmail(email)

	if err != nil {
		return 0, err
	}

	return data.Id, nil
}
