package services

import "github.com/OlegDjur/API-MongoDB-Redis/models"

type UserService interface {
	FindUserByID(string) (*models.DBResponse, error)
	FindUserByEmail(string) (*models.DBResponse, error)
}
