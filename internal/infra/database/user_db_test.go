package database

import (
	"testing"

	"github.com/gusgd/apigo/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db := SetupDatabase(t, &entity.User{})
	user := CreateTestUser(db)
	userDB := NewUser(db)
	err := userDB.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	AssertUserEqual(t, user, &userFound)
}

func TestFindByEmail(t *testing.T) {
	db := SetupDatabase(t, &entity.User{})
	user := CreateTestUser(db)
	userDB := NewUser(db)
	err := userDB.Create(user)
	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	AssertUserEqual(t, user, userFound)
}
