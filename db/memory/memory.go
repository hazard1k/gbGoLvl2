package memory

import (
	"golvl2/app/domain/db/environments"
	"golvl2/app/domain/db/users"
	"sync"
)

type MemStore struct {
	users users.Repository
	envs  environments.Repository
}

func (m MemStore) Users() users.Repository {
	return m.users
}

func (m MemStore) Environment() environments.Repository {
	return m.envs
}

func NewStore() *MemStore {
	return &MemStore{
		users: &usersRepository{users: &sync.Map{}},
		envs:  &environmentsRepository{},
	}
}
