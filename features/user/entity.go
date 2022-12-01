package user

import (
	"time"
)

type Core struct {
	ID        uint
	Name      string `validate:"required"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required"`
	Phone     string `validate:"required"`
	Address   string
	Role      string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ServiceInterface interface {
	GetAlluser() (data []Core, err error)
	InsertUser(input Core) error
	GetById(id int) (Core, error)
	Delete(id int) error
	UpdateUser(id int, input Core) error
}

type RepositoryInterface interface {
	GetAlluser() ([]Core, error)
	InsertUser(input Core) error
	GetById(id int) (Core, error)
	Delete(id int) error
	UpdateUser(id int, input Core) error
}
