package order

import (
	"fmt"
	"time"

	"order-pack/internal/product"

	"gorm.io/gorm"
)

type Item struct {
	Count   int             `json:"count"`
	Product product.Product `json:"product"`
}

type Order struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ItemCount int    `json:"item_count"`
	Items     []Item `json:"items"`
}

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) Create(o *Order) error {
	if err := s.db.Create(o).Error; err != nil {
		return fmt.Errorf("could not create order, %w", err)
	}

	return nil
}

// FulfillOrder ...
func FulfillOrder(order *Order, products []product.Product) *Order {
	remainingItems := order.ItemCount

	for i, curr := range products {
		if remainingItems <= 0 {
			// order has already fulfilled
			break
		}

		if i != len(products)-1 {
			next := products[i+1]

			if i < len(products)-1 && remainingItems > curr.PackageSize && remainingItems < next.PackageSize {
				packsNeeded := (remainingItems + next.PackageSize - 1) / next.PackageSize
				order.Items = append(order.Items, Item{
					Count:   packsNeeded,
					Product: next,
				})
				remainingItems -= packsNeeded * next.PackageSize
			}
		}

		packsNeeded := remainingItems / curr.PackageSize
		remainingItems -= packsNeeded * curr.PackageSize
		if packsNeeded > 0 {
			order.Items = append(order.Items, Item{
				Count:   packsNeeded,
				Product: curr,
			})
		}
	}

	if remainingItems > 0 {
		order.Items = append(order.Items, Item{
			Count:   1,
			Product: products[len(products)-1],
		})
	}

	return order
}
