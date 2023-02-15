package services

import "github.com/OlegDjur/API-MongoDB-Redis/models"

type UserService interface {
	FindUserByID(id string) (*models.DBResponse, error)
	FindUserByEmail(email string) (*models.DBResponse, error)
}
