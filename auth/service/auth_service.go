package service

import (
	"context"
	"efishery-auth/entity/form"
	"efishery-auth/entity/model"
	"efishery-auth/errorlib"
	"fmt"
	"time"

	"crypto/md5"

	"github.com/aidarkhanov/nanoid"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthService interface {
	Register(form.RegisterForm, context.Context) (*string, error)
	Login(form.LoginForm, context.Context) (*string, error)
	Explain(string, context.Context) (*model.Claim, error)
}

type authService struct {
	Repo      rel.Repository
	JwtSecret string
}

func NewAuthService(repo rel.Repository, jwtSecret string) AuthService {
	return &authService{
		Repo:      repo,
		JwtSecret: jwtSecret,
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

func (s *authService) Login(form form.LoginForm, ctx context.Context) (*string, error) {
	var user model.User
	err := s.Repo.Find(ctx, &user, where.Eq("phone", form.Phone))

	if err != nil {
		return nil, errorlib.ErrLogin
	}

	hashedPassword := fmt.Sprintf("%x", md5.Sum([]byte(form.Password)))
	if user.Password != hashedPassword {
		return nil, errorlib.ErrLogin
	}

	claims := model.Claim{
		Username: user.Username,
		Phone:    user.Phone,
		Role:     user.Role,
	}
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    "efishery-auth",
		Subject:   user.Username,
		ID:        uuid.New().String(),
		Audience:  []string{"efishery-fetch"},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.JwtSecret))
	if err != nil {
		return nil, errorlib.ErrSomethingWrong
	}

	return &tokenString, nil
}

func (s *authService) Explain(JwtToken string, ctx context.Context) (*model.Claim, error) {
	token, err := jwt.ParseWithClaims(JwtToken, &model.Claim{}, s.validateToken)

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*model.Claim); ok && token.Valid {
		return claims, nil
	}

	return nil, errorlib.ErrInvalidJwt
}

func (s *authService) validateToken(token *jwt.Token) (interface{}, error) {
	//this will validate the signing method
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errorlib.ErrInvalidJwt
	}

	//returning jwt secret
	return []byte(s.JwtSecret), nil
}
