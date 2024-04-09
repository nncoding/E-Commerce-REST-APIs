package middlewarerepositories

import "github.com/jmoiron/sqlx"

type IMiddlewareRepository interface {
}

type middlewarerepository struct {
	db *sqlx.DB
}

func MiddlewaresRepository(db *sqlx.DB) IMiddlewareRepository {
	return &middlewarerepository{
		db: db,
	}
}
