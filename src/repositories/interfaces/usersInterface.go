package interfaces

import (
	"blog/src/models"
)

type IUser interface {
	Get(id int) (models.Users, error)
	GetByEmail(email string) (models.Users, error)
	GetAll() ([]models.Users, error)
	Create(email string, password []byte) error
	Edit(id int, email string, password []byte) error
	Delete(id int) error
}
