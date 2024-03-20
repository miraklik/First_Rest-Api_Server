package teststore

import (
	"golangApi/https-rest-api/internal/app/model"
	"golangApi/https-rest-api/internal/app/store"
)

type Store struct {
	userRepository *UserRepository
}

func New() *Store {
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		user:  make(map[int]*model.User),
	}

	return s.userRepository
}
