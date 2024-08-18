package impl

import (
	"contact-management-restful/exception"
	"contact-management-restful/helper"
	"contact-management-restful/models/domains"
	"contact-management-restful/models/dto"
	"contact-management-restful/repositories/contracts"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository contracts.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository contracts.UserRepository, DB *sql.DB, validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{UserRepository: userRepository, DB: DB, Validate: validate}
}

func (s *UserServiceImpl) Register(ctx context.Context, req dto.RegisterRequest) *dto.RegisterResponse {
	err := s.Validate.Struct(req)
	if err != nil {
		panic(err)
	}

	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	user, err := s.UserRepository.FindByUsername(ctx, tx, req.Username)
	if err == nil {
		panic(exception.NewUnauthorized("Username already registered"))
	}

	token := "123jnjndasjkn123jnjska"
	hashedPassword, err := helper.HashedPassword(req.Password)
	if err != nil {
		panic(err)
	}

	user = domains.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Name:     req.Name,
		Token:    token,
	}

	user = s.UserRepository.Save(ctx, tx, user)

	return helper.RegisterResponse(&user)
}

func (s *UserServiceImpl) Login(ctx context.Context, req dto.LoginRequest) *dto.LoginResponse {
	err := s.Validate.Struct(req)
	if err != nil {
		panic(err)
	}

	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	user, err := s.UserRepository.FindByUsername(ctx, tx, req.Username)
	if err != nil {
		panic(exception.NewUnauthorized("wrong credentials"))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		panic(exception.NewUnauthorized("wrong credentials"))
	}

	token := "njkn213nkasndkn"
	user.Token = token
	user = s.UserRepository.Update(ctx, tx, user)

	loginResponse := dto.LoginResponse{
		Name:     user.Name,
		Username: user.Username,
		Token:    user.Token,
	}

	return &loginResponse
}

func (s *UserServiceImpl) Update(ctx context.Context, req dto.UpdateUserRequest) *dto.UpdateUserResponse {
	err := s.Validate.Struct(req)
	if err != nil {
		panic(err)
	}

	user := ctx.Value("user").(domains.User)
	user.Username = req.Username
	user.Name = req.Name
	hashedPassword, err := helper.HashedPassword(req.Password)
	if err != nil {
		panic(err)
	}
	user.Password = string(hashedPassword)

	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	user = s.UserRepository.Update(ctx, tx, user)

	updateResponse := dto.UpdateUserResponse{
		Name:     user.Name,
		Username: user.Username,
		Token:    user.Token,
	}

	return &updateResponse
}

func (s *UserServiceImpl) Logout(ctx context.Context) {
	user := ctx.Value("user").(domains.User)
	user.Token = ""

	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	_ = s.UserRepository.Update(ctx, tx, user)
}
