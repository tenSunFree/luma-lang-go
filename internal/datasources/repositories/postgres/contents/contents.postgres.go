package postgres

import (
	"github.com/jmoiron/sqlx"
	repointerface "github.com/snykk/go-rest-boilerplate/internal/datasources/repositories/interface"
)

type postgreContentRepository struct {
	conn *sqlx.DB
}

func NewContentRepository(conn *sqlx.DB) repointerface.ContentRepository {
	return &postgreContentRepository{conn: conn}
}
