package database

import "github.com/gusgd/apigo/internal/entity"

type UserDBInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ProductDBInterface interface {
	Create(product *entity.Product) error
	FindAllProduct(page, limit int, sort string) ([]entity.Product, error)
	FindByProductName(productName string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
