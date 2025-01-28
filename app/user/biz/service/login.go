package service

import (
	"context"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
	"zqzqsb.com/gomall/app/user/biz/dal/mysql"
	"zqzqsb.com/gomall/app/user/biz/model"
	user "zqzqsb.com/gomall/app/user/kitex_gen/user"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.

	log.Println(req.String())
	// verify empty
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("empty Email or Password")
	}

	// fetch user
	row, err := model.GetbyEmail(mysql.DB, req.Email)
	if err != nil {
		return nil, err
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHashed), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	// return user id
	resp = &user.LoginResp{
		UserId: int32(row.ID),
	}

	return
}
