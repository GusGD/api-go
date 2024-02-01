package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/gusgd/apigo/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewProduct(t *testing.T) {
	db := SetupDatabase(t, &entity.Product{})
	product, err := entity.NewProduct("testProduct", 10.98)
	assert.NoError(t, err)

	productDB := NewProduct(db)
	err = productDB.Create(product)

	assert.Nil(t, err)
	assert.NotEmpty(t, product.ID)
	assert.NotEmpty(t, product.Name)
	assert.NotEmpty(t, product.Price)
}

func TestFindAllProduct(t *testing.T) {
	db := SetupDatabase(t, &entity.Product{})

	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(product)
	}
	productDB := NewProduct(db)

	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)
}

func TestFindByIdProduct(t *testing.T) {
	db := SetupDatabase(t, &entity.Product{})
	product := CreateRandomProduct(db)
	productDB := NewProduct(db)
	product, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.NotEmpty(t, product.Name)
}
func TestFindByNameProduct(t *testing.T) {
	db := SetupDatabase(t, &entity.Product{})
	product := CreateRandomProduct(db)
	productDB := NewProduct(db)
	product, err := productDB.FindByName(product.Name)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.Name)
}

func TestUpdateProduct(t *testing.T) {
	db := SetupDatabase(t, &entity.Product{})
	product := CreateRandomProduct(db)
	productDB := NewProduct(db)
	product.Name = "pRODUCT"
	err := productDB.Update(product)
	assert.NoError(t, err)
	product, err = productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "pRODUCT", product.Name)
}

func TestDeleteProduct(t *testing.T) {
	db := SetupDatabase(t, &entity.Product{})
	product := CreateRandomProduct(db)
	productDB := NewProduct(db)
	err := productDB.Delete(product)
	assert.NoError(t, err)
	product, err = productDB.FindByID(product.ID.String())
	assert.Error(t, err)
	assert.Empty(t, product)
}
