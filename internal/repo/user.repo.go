package repo

import (
	"github.com/ntquang/ecommerce/global"
	"github.com/ntquang/ecommerce/internal/model"
)

// type UserRepo struct{}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// func (ur *UserRepo) GetInfoUser() string {
// 	return "quangdz"
// }

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {
}

func (ur *userRepository) GetUserByEmail(email string) bool {
	// Implement the GetUserByEmail method here
	row := global.Pdb.Table("go_crm_user").Where("usr_email = ?", email).First(&model.GoCrmUser{}).RowsAffected
	return row != NumberNull
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
