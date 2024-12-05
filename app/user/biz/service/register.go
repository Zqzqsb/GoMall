package service

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"zqzqsb.com/gomall/app/user/biz/dal/mysql"
	"zqzqsb.com/gomall/app/user/biz/model"
	user "zqzqsb.com/gomall/app/user/kitex_gen/user"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.

	// verify empty
	if req.Password == "" || req.Email == "" || req.PasswordConfirm == "" {
		return nil, errors.New("empty Email or Password")
	}

	// verify consistency
	if req.Password != req.PasswordConfirm {
		return nil, errors.New("password not match")
	}
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(passwordHashed),
	}

	err = model.Create(mysql.DB, newUser)
	if err != nil {
		return nil, err
	}

	return &user.RegisterResp{UserId: int32(newUser.ID)}, nil
}
