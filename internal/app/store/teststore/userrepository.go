package teststore

import (
	"golangApi/https-rest-api/internal/app/model"
	"golangApi/https-rest-api/internal/app/store"
)

type UserRepository struct {
	store *Store
	user  map[int]*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	u.ID = len(r.user) + 1
	r.user[u.ID] = u

	return nil
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	u, ok := r.user[email]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	for _, u := range r.user {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, store.ErrRecordNotFound
}
