package repository

import "tutorial/entities"

type AdminRepository interface {
	Creating(adminEntity *entities.Admin) (*entities.Admin, error)
	FindByID(adminId string) (*entities.Admin, error)
}
