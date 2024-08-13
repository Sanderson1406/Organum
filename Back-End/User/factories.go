package factories

import (
	"idgen"
	"users"
)

type UserFactory struct{}

func (f *UserFactory) CreateUser(name, email, password string) *users.User {
	id := idgen.GenerateID()
	metadata := users.NewMetadata()
	return users.NewUser(name, email, password, nil, metadata)
}

func (f *UserFactory) CreateUserWithMetadata(name, email, password string, m *Metadada) *users.User {
	id := idgen.GenerateID()
	return users.NewUser(name, email, password, nil, m)
}
