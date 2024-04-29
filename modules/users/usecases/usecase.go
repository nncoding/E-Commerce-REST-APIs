package userUsecases

import (
	"github.com/nncoding/go-basic/config"
	"github.com/nncoding/go-basic/modules/users"
	userRepositories "github.com/nncoding/go-basic/modules/users/repositories"
)

type IUsersUsecase interface {
	InsertCustomer(req *users.UserRegisterReq) (*users.UserPassport, error)
}

type usersUsecases struct {
	cfg              config.IConfig
	userRepositories userRepositories.IUserRepository
}

func UsersUsecases(cfg config.IConfig, usersRepository userRepositories.IUserRepository) IUsersUsecase {
	return &usersUsecases{
		cfg:              cfg,
		userRepositories: usersRepository,
	}
}

func (u *usersUsecases) InsertCustomer(req *users.UserRegisterReq) (*users.UserPassport, error) {
	// Hashing password
	if err := req.BcryptHashing(); err != nil {
		return nil, err
	}

	// Insert User
	result, err := u.userRepositories.InsertUser(req, false)
	if err != nil {
		return nil, err
	}

	return result, nil
}
