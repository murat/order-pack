package pack

import (
	"order-pack/internal/database"

	"gorm.io/gorm"
)

type Pack struct {
	gorm.Model
	PackageSize int `json:"size"`
}

type Service struct {
	db *database.Database
}

func NewService(db *database.Database) Service {
	return Service{db: db}
}

func (s Service) Get() ([]Pack, error) {
	return []Pack{
		{
			PackageSize: 100,
		},
	}, nil
}

func (s Service) Find() (Pack, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) Create(p Pack) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) Update(p Pack) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) Delete(p Pack) error {
	//TODO implement me
	panic("implement me")
}
