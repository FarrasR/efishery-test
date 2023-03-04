package service

import (
	"context"
	"efishery-auth/entity/form"
	"efishery-auth/entity/model"
	"efishery-auth/errorlib"
	"fmt"

	"crypto/md5"

	"github.com/aidarkhanov/nanoid"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
)

type AuthService interface {
	Register(form.RegisterForm, context.Context) (*string, error)
	Login(form.LoginForm, context.Context) (*string, error)
	Explain(string, context.Context) (any, error)
}

type authService struct {
	Repo rel.Repository
}

func NewAuthService(repo rel.Repository) AuthService {
	return &authService{
		Repo: repo,
	}
}

func (s *authService) Register(form form.RegisterForm, ctx context.Context) (*string, error) {

	var user model.User

	err := s.Repo.Find(ctx, &user, where.Eq("phone", form.Phone))

	if err != rel.ErrNotFound {
		return nil, errorlib.ErrPhoneNumberRegistered
	}

	rawPassword, _ := nanoid.Generate(nanoid.DefaultAlphabet, 4)

	//this could use a better algorithm
	password := fmt.Sprintf("%x", md5.Sum([]byte(rawPassword)))

	newUser := model.User{
		Username: form.Username,
		Phone:    form.Phone,
		Role:     form.Role,
		Password: password,
	}

	err = s.Repo.Insert(ctx, &newUser)

	if err != nil {
		return nil, errorlib.ErrSomethingWrong
	}

	return &rawPassword, nil
}

func (s *authService) Login(form.LoginForm, context.Context) (*string, error) {
	return nil, nil
}

func (s *authService) Explain(string, context.Context) (any, error) {
	return nil, nil
}
