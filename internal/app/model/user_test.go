package model_test

import (
	"golangApi/https-rest-api/internal/app/model"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestUser_Validate(t *testing.T){
	testCases := []struct { 
		name string
		u func() *model.User
		isValid bool
	}{
		{
			name: "Valid",
			u: func() *model.User{
				return model.TestUser()
			},
			isValid: true,
		},
		{
			name: "empty email",
			u: func() *model.User{
				u := model.TestUser(t)
				u.Email = ""

				return u
			},
			isValid: false,
		},
		{
			name: "invalid email",
			u: func() *model.User{
				u := model.TestUser(t)
				u.Email = "invalid"

				return u
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *model.User{
				u := model.TestUser(t)
				u.Password = ""

				return u
			},
			isValid: false,
		},
		{
			name: "short password",
			u: func() *model.User{
				u := model.TestUser(t)
				u.Password = "short"

				return u
			},
			isValid: false,
		},
		name: "with encrypto password",
			u: func() *model.User{
				u := model.TestUser(t)
				u.Password = ""
				u.EncryptoPassword = "with encrypto password"

				return u
			},
			isValid: false,
		},
	}
	
	for _, tc := range testCases{
		t.Run(tc.name, func(t *testing.T){
			if tc.isValid{
				assert.NoError(t, tc.u().Validate())
			}else{
				assert.Error(t, tc.u().Validate())
			}
		})
	}



func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptoPassword)
}
