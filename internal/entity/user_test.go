package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	usr, err := NewUser("test", "test@test.com", "password")
	assert.Nil(t, err)
	assert.NotNil(t, usr)
	assert.NotEmpty(t, usr.ID)
	assert.NotEmpty(t, usr.Password)
	assert.Equal(t, "test", usr.Name)
	assert.Equal(t, "test@test.com", usr.Email)

}

func TestUser_ValidatePassword(t *testing.T) {
	usr, err := NewUser("test", "test@test.com", "password")
	assert.Nil(t, err)
	assert.True(t, usr.ValidatePassword("password"))
	assert.False(t, usr.ValidatePassword("password123"))
	assert.NotEqual(t, "password", usr.Password)
}
