package repo

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
	return false
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
