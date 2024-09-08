package services

import (
	"github.com/ntquang/ecommerce/internal/repo"
	"github.com/ntquang/ecommerce/response"
)

// type UserService struct {
// 	userRepo *repo.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepo: repo.NewUserRepo(),
// 	}
// }

// func (us *UserService) GetInfoUser() string {
// 	return us.userRepo.GetInfoUser()
// }

// INTERFACE VERSION
type IUserservice interface {
	Register(email string, purpose string) int
}

type userService struct { //dua ve private vi IUserservice se truc tiep lam viec voi controller
	userRepo repo.IUserRepository
}

func NewUserService(
	userRepo repo.IUserRepository,
) IUserservice {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) Register(email string, purpose string) int {
	// 1. check email exists
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrEmailAlreadyExists
	}
	return response.ErrCodeSuccess
}
