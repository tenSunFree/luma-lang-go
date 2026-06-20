package lives

import "github.com/jmoiron/sqlx"

type postgreLiveRepository struct {
	conn *sqlx.DB
}

func NewPostgreLiveRepository(conn *sqlx.DB) *postgreLiveRepository {
	return &postgreLiveRepository{conn: conn}
}
