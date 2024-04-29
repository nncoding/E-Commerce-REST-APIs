package userRepositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/nncoding/go-basic/modules/users"
	userPatterns "github.com/nncoding/go-basic/modules/users/patterns"
)

type IUserRepository interface {
	InsertUser(req *users.UserRegisterReq, isAdmin bool) (*users.UserPassport, error)
}

type userRepositories struct {
	db *sqlx.DB
}

func UserRepositories(db *sqlx.DB) IUserRepository {
	return &userRepositories{
		db: db,
	}
}

func (r *userRepositories) InsertUser(req *users.UserRegisterReq, isAdmin bool) (*users.UserPassport, error) {
	result := userPatterns.InsertUser(r.db, req, isAdmin)

	var err error
	if isAdmin {
		result, err = result.Admin()
		if err != nil {
			return nil, err
		}
	} else {
		result, err = result.Customer()
		if err != nil {
			return nil, err
		}
	}

	// Get  Result From Inserting
	user, err := result.Result()
	if err != nil {
		return nil, err
	}

	return user, nil
}
