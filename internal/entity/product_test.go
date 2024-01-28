package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("test", 10.98)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "test", p.Name)
	assert.Equal(t, 10.98, p.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 10.98)
	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenNameIsValid(t *testing.T) {
	p, err := NewProduct("test", 10.98)
	assert.Nil(t, p)
	assert.Equal(t, ErrInvalidName, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Product10", 0)
	assert.Nil(t, p)
	assert.Equal(t, ErrPriceIsRequired, err)
}
func TestProductWhenPriceIsValid(t *testing.T) {
	p, err := NewProduct("Product10", -90.90)
	assert.Nil(t, p)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProductValidate(t *testing.T) {
	p, err := NewProduct("Product2", 89.09)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
