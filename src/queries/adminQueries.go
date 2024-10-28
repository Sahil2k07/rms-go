package queries

import "database/sql"

type AdminQueries struct {
	db *sql.DB
}

func NewAdminQueries(db *sql.DB) *AdminQueries {
	return &AdminQueries{db: db}
}
