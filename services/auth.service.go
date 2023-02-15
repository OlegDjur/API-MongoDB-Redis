package services

import "github.com/OlegDjur/API-MongoDB-Redis/models"

type AuthService interface {
	SignUpUser(*models.SignUpInput) (*models.DBResponse, error)
	SignInInput(*models.SignInInput) (*models.DBResponse, error)
}
