package service

import (
	"final-project-enigma/dto/request"
	"final-project-enigma/dto/response"
)

type AuthService interface {
	RegisterAccount(req request.RegisterAccountRequest) (resp response.RegisterAccountResponse, err error)
}