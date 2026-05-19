package service

import (
	"ai-workspace-backend/dao"
	"ai-workspace-backend/model"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Email string `json:"email"`
	Username string	`json:"username"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	ID uint `json:"id"`
	Email string `json:"email"`
	Username string `json:"username"`
}

func Register(req RegisterRequest) (*RegisterResponse, error) {
	if req.Email == "" || req.Username == "" || req.Password == "" {
		return nil, errors.New("email, username and password are required")
	}

	existingUser, err := dao.GetUserByEmail(req.Email)
	if err == nil && existingUser != nil {
		return nil, errors.New("email already registered")
	}

	if err != nil && !dao.IsRecordNotFound(err) {
		return nil, fmt.Errorf("check existing user: %w", err)
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	user := &model.User{
		Email: req.Email,
		Username: req.Username,
		Password: string(passwordHash),
	}

	if err := dao.CreateUser(user); err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return &RegisterResponse{
		ID: user.ID,
		Email: user.Email,
		Username: user.Username,
	},nil


}
