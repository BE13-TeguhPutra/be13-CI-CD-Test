package book

import (
	"time"
)

type Core struct {
	ID        uint
	Book      string
	Publisher string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ServiceInterface interface {
	GetBook() error
}

type RepositoryInterface interface {
	GetBook() error
}
