package teststore_test

import (
	"golangApi/https-rest-api/internal/app/model"
	"golangApi/https-rest-api/internal/app/store"
	"golangApi/https-rest-api/internal/app/store/teststore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepositore_Create(t *testing.T) {

	s := teststore.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepositore_Find(t *testing.T) {
	s := teststore.New()
	u1 := model.TestUser(t)
	u1.Email = email
	s.User().Create(u1)

	u2, err := s.User().Find(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)

}

func TestUserRepositore_FindByEmail(t *testing.T) {

	s := teststore.New(db)
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)

	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)

}
