package service

import (
	"context"
	"errors"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"golang.org/x/crypto/bcrypt"
	"zqzqsb.com/gomall/app/user/biz/dal/mysql"
	"zqzqsb.com/gomall/app/user/biz/model"
	user "zqzqsb.com/gomall/app/user/kitex_gen/user"
)

var (
	ErrEmptyFields      = errors.New("邮箱或密码不能为空")
	ErrInvalidEmail     = errors.New("无效的邮箱格式")
	ErrEmailExists      = errors.New("邮箱已被注册")
	ErrPasswordMismatch = errors.New("两次输入的密码不匹配")
	ErrPasswordTooWeak  = errors.New("密码强度不够（至少8位，包含大小写字母、数字和特殊字符）")
	emailRegex          = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

type RegisterService struct {
	ctx context.Context
}

func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

func (s *RegisterService) validateEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func (s *RegisterService) validatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSpecial
}

func (s *RegisterService) checkEmailExists(email string) bool {
	_, err := model.GetbyEmail(mysql.DB, email)
	return err == nil
}

func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// 1. 基本字段验证
	if req.Password == "" || req.Email == "" || req.PasswordConfirm == "" {
		return nil, ErrEmptyFields
	}

	// 2. 邮箱格式验证
	if !s.validateEmail(req.Email) {
		return nil, ErrInvalidEmail
	}

	// 3. 邮箱唯一性检查
	if s.checkEmailExists(req.Email) {
		return nil, ErrEmailExists
	}

	// 4. 密码匹配验证
	if req.Password != req.PasswordConfirm {
		return nil, ErrPasswordMismatch
	}

	// 5. 密码强度验证
	if !s.validatePassword(req.Password) {
		return nil, ErrPasswordTooWeak
	}

	// 6. 密码加密
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		hlog.CtxErrorf(s.ctx, "Failed to hash password: %v", err)
		return nil, errors.New("服务器内部错误")
	}

	// 7. 创建用户
	newUser := &model.User{
		Email:             strings.ToLower(req.Email), // 统一转换为小写
		PasswordHashed:    string(passwordHashed),
		PasswordChangedAt: time.Now(),
	}

	err = model.Create(mysql.DB, newUser)
	if err != nil {
		hlog.CtxErrorf(s.ctx, "Failed to create user: %v", err)
		return nil, errors.New("用户创建失败")
	}

	hlog.CtxInfof(s.ctx, "User registered successfully with email: %s", req.Email)
	return &user.RegisterResp{UserId: int32(newUser.ID)}, nil
}
