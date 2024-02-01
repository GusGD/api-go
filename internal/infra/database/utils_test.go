package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/gusgd/apigo/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateTestUser(db *gorm.DB) *entity.User {
	user, err := entity.NewUser("test", "test@email.com", "testsenha")
	if err != nil {
		panic(err)
	}
	return user
}

func AssertUserEqual(t *testing.T, user *entity.User, userFound *entity.User) {
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotEmpty(t, user.Password)
	assert.NotNil(t, user.Password)
}

func SetupDatabase(t *testing.T, model interface{}) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(model)
	return db
}

func SetupDatabaseUser(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	return db
}

func CreateRandomProduct(db *gorm.DB) *entity.Product {
	product, err := entity.NewProduct(fmt.Sprintf("Product %d", rand.Intn(100)), rand.Float64()*100)
	if err != nil {
		panic(err)
	}
	db.Create(product)
	return product
}
