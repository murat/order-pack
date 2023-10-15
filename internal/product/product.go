package product

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string `json:"name"`
	PackageSize int    `json:"package_size"`
}

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) Get() ([]Product, error) {
	var products []Product
	if err := s.db.Find(&products).Error; err != nil {
		return nil, fmt.Errorf("could not fetch products, %w", err)
	}

	return products, nil
}

func (s *Service) GetSortedBy(sort string) ([]Product, error) {
	var products []Product
	if err := s.db.Order(sort).Find(&products).Error; err != nil {
		return nil, fmt.Errorf("could not fetch products, %w", err)
	}

	return products, nil
}

func (s *Service) Find(id uint) (*Product, error) {
	var product Product
	if err := s.db.First(&product, id).Error; err != nil {
		return nil, fmt.Errorf("could not fetch products, %w", err)
	}

	return &product, nil
}

func (s *Service) Create(p *Product) error {
	if err := s.db.Create(p).Error; err != nil {
		return fmt.Errorf("could not create product, %w", err)
	}

	return nil
}

func (s *Service) Update(_ *Product) error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Delete(_ *Product) error {
	//TODO implement me
	panic("implement me")
}
