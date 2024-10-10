package imple

import (
	"context"
	"fmt"

	"github.com/ntquang/ecommerce/internal/database"
)

type sUserLogin struct {
	// implement the IUserLogin interface here
	r *database.Queries
}

func NewUserLoginImpl(r *database.Queries) *sUserLogin {
	return &sUserLogin{
		r: r,
	}
}

// implement the IUserLogin interface here

func (s *sUserLogin) Login(ctx context.Context) error {
	fmt.Println("login already work")
	return nil
}

func (s *sUserLogin) Register(ctx context.Context) error {
	return nil
}

func (s *sUserLogin) VerifyOtp(ctx context.Context) error {
	return nil
}

func (s *sUserLogin) UpdatePaswordRegister(ctx context.Context) error {
	return nil
}
