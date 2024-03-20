package sqlstore

import (
	"database/sql"
	"golangApi/https-rest-api/internal/app/model"
	"golangApi/https-rest-api/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}
	return r.store.db.QueryRow("INSERT INTO users (email, enctcrypto_password) VALUES ($1, $2) RETURNING id", u.Email, u.EncryptoPassword).Scan(&u.ID)
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow("SELECT id, email,  encrypto_password FROM users WHERE id = $1", id).Scan(&u.ID, &u.Email, &u.EncryptoPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return nil, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow("SELECT id, email,  encrypto_password FROM users WHERE email = $1", email).Scan(&u.ID, &u.Email, &u.EncryptoPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return nil, nil
}
