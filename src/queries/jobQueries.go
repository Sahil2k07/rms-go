package queries

import "database/sql"

type JobQueries struct {
	db *sql.DB
}

func NewJobQueries(db *sql.DB) *JobQueries {
	return &JobQueries{db: db}
}
