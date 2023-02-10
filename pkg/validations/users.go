package validations

import (
	simplebankpb "github.com/orlandorode97/simple-bank/generated/simplebank"
)

type CreateUserValidator struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
	FullName string `validate:"required"`
	Email    string `validate:"required,email"`
}

func NewCreateUserValidator(req *simplebankpb.CreateUserRequest) *CreateUserValidator {
	return &CreateUserValidator{
		Username: req.Username,
		Password: req.Password,
		FullName: req.FullName,
		Email:    req.Email,
	}
}

type LoginValidator struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

func NewLoginValidator(req *simplebankpb.LoginRequest) *LoginValidator {
	return &LoginValidator{
		Username: req.Username,
		Password: req.Password,
	}
}