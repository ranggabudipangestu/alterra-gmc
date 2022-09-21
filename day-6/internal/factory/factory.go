package factory

import (
	"github.com/ranggabudipangestu/hexagonal-api/database"
	"github.com/ranggabudipangestu/hexagonal-api/internal/repository"
)

type Factory struct {
	UserRepository repository.User
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		repository.NewUser(db),
	}
}
