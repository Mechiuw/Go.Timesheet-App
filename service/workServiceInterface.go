package service

import (
	"final-project-enigma/dto/request"
	"final-project-enigma/dto/response"
)

type WorkService interface {
	CreateWork(request request.WorkRequest) (response.WorkResponse, error)
	UpdateWork(id string, request request.WorkRequest) (response.WorkResponse, error)
	DeleteWork(id string) error
}